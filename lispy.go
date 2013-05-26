// Simple, Elegant yet Extendible Syntax System. (p:Hello World in Paragraph!)
package lispy

import (
	"bytes"
	_html "html"
	html "html/template"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// (name:value), (name:(name:value)), (name:content|key:value|key:value)

var starterRegExp = regexp.MustCompile(`\(([\p{L}\p{N}-_]+?):`)

const (
	openRune  = rune('(')
	closeRune = rune(')')
)

type lispyMap map[string][]LispyHandler

func (li lispyMap) Set(name string, function LispyFunc) {
	li.SetHandler(name, function)
}

func (li lispyMap) SetHandler(name string, lispyhandler LispyHandler) {
	name = strings.TrimSpace(strings.ToLower(name))
	if len(li[name]) <= 0 {
		li[name] = append(li[name], lispyhandler)
	} else {
		li[name][0] = lispyhandler
	}
}

// Lispy Map
type LispyMap map[string]LispyFunc

// Setter
func (li LispyMap) Set(name string, function func(li *Lispy) string) {
	li[name] = function
}

// Lispy Handler Interface
type LispyHandler interface {
	Lispy(*Lispy) string
}

// Lispy Handler Map
type LispyHandlerMap map[string]LispyHandler

// Setter
func (li LispyHandlerMap) Set(name string, handle LispyHandler) {
	li[name] = handle
}

type LispyFunc func(*Lispy) string

func (li LispyFunc) Lispy(lii *Lispy) string {
	return li(lii)
}

// Lispy Structure
type Lispy struct {
	sync.RWMutex
	Name          string
	Content       string
	htmlEscape    bool
	restrictParam bool
	allowedNames  []string
	param         map[string][]string
	code          lispyMap
	first         bool
	linebreak     bool
	paramParsed   bool
}

// Construct with Default Map
func New() *Lispy {
	li := &Lispy{
		Name:          "",
		Content:       "",
		htmlEscape:    true,
		restrictParam: false,
		allowedNames:  []string{},
		param:         map[string][]string{},
		code:          lispyMap{},
		first:         true,
		linebreak:     false,
		paramParsed:   false,
	}

	DefaultMap.RLock()

	li.AddFuncMap(DefaultMap.Map).AddHandlerMap(DefaultMap.HandlerMap)

	DefaultMap.RUnlock()

	return li
}

// Copy
func (li *Lispy) Copy() *Lispy {
	li.RLock()
	defer li.RUnlock()

	allowedNames := append([]string{}, li.allowedNames...)

	code := lispyMap{}

	for key, function := range li.code {
		code[key] = append(code[key], function...)
	}

	newli := &Lispy{}
	*newli = *li
	newli.code = code
	newli.allowedNames = allowedNames
	newli.RWMutex = sync.RWMutex{}
	newli.param = map[string][]string{}

	return newli
}

// Set Function
func (li *Lispy) SetFunc(name string, function LispyFunc) *Lispy {
	li.RLock()
	defer li.RUnlock()
	li.code.Set(name, function)
	return li
}

// Set Handler
func (li *Lispy) SetHandler(name string, lispyhandler LispyHandler) *Lispy {
	li.RLock()
	defer li.RUnlock()
	li.code.SetHandler(name, lispyhandler)
	return li
}

// Add Function Map
func (li *Lispy) AddFuncMap(lispymap LispyMap) *Lispy {
	if lispymap == nil {
		return li
	}

	for name, function := range lispymap {
		li.SetFunc(name, function)
	}
	return li
}

// Add Handler Map
func (li *Lispy) AddHandlerMap(lispyhandlermap LispyHandlerMap) *Lispy {
	if lispyhandlermap == nil {
		return li
	}

	for name, lispyhandler := range lispyhandlermap {
		li.SetHandler(name, lispyhandler)
	}
	return li
}

// Make string as safe!
func (li *Lispy) Safe(str string) html.HTML {
	return html.HTML(str)
}

func (li *Lispy) nameAllowed() bool {
	if len(li.allowedNames) <= 0 {
		return true
	}

	for _, name := range li.allowedNames {
		if name == li.Name {
			return true
		}
	}

	return false
}

type previousState struct {
	htmlEscape, first, linebreak bool
}

type openclose struct {
	opener bool
	pos    int
}

func (op openclose) Pos() int {
	return op.pos
}

type openclosers []openclose

func (op openclosers) Len() int {
	return len(op)
}

func (op openclosers) Less(i, j int) bool {
	return op[i].pos < op[j].pos
}

func (op openclosers) Swap(i, j int) {
	op[i], op[j] = op[j], op[i]
}

// Alais of Render, but marks the string as safe!
func (li *Lispy) RenderSafe(str string) html.HTML {
	return li.Safe(li.Render(str))
}

// Render Syntax from String, returns rendered String
func (li *Lispy) Render(str string) string {
	li = li.Copy()
	li.SetFunc("br", Br)

	if li.code == nil {
		return str
	}

	previous := previousState{li.htmlEscape, li.first, li.linebreak}

	if li.htmlEscape {
		str = html.HTMLEscapeString(str)
		li.htmlEscape = false
	}

	if li.linebreak {
		str = strings.Replace(str, "\r\n", "<br />", -1)
		str = strings.Replace(str, "\r", "<br />", -1)
		str = strings.Replace(str, "\n", "<br />", -1)
		li.linebreak = false
	}

	if li.first {
		str = strings.Replace(str, `\|`, "&#124;", -1)
		str = strings.Replace(str, `\:`, "&#58;", -1)
		str = strings.Replace(str, `\(`, "&#40;", -1)
		str = strings.Replace(str, `\)`, "&#41;", -1)
		li.first = false
	}

	processedStr := ""

	for {
		index := starterRegExp.FindStringIndex(str)
		if index == nil {
			processedStr += str
			break
		}

		pos := index[0]

		submatch := starterRegExp.FindStringSubmatch(str)
		lenght := len(submatch[0])
		li.Name = strings.ToLower(submatch[1])

		content_pos := pos + lenght
		li.Content = str[content_pos:]

		openers_indexes := indexRune(li.Content, openRune)
		closers_indexes := indexRune(li.Content, closeRune)

		content_lenght := len(li.Content)
		open := 1

		openers := openclosers{}

		for _, opener := range openers_indexes {
			openers = append(openers, openclose{true, opener})
		}

		for _, closer := range closers_indexes {
			openers = append(openers, openclose{false, closer})
		}

		sort.Sort(openers)

		// Find closer for current opener!
		for _, opener := range openers {
			if opener.opener {
				open++
			} else {
				open--
			}
			if open <= 0 {
				content_lenght = opener.pos
				break
			}
		}

		li.paramParsed = false

		li.Content = li.Content[:content_lenght]

		if len(li.code[li.Name]) <= 0 || !li.nameAllowed() {
			processedStr += str[:pos] + li.Content
		} else {
			processedStr += str[:pos] + li.code[li.Name][0].Lispy(li)
			li.param = map[string][]string{}
		}

		// Reset state
		li.Name = ""
		li.Content = ""

		if len(str) >= content_pos+content_lenght+1 {
			str = str[content_pos+content_lenght+1:]
		} else {
			str = str[content_pos+content_lenght:]
		}
	}

	// Restore previous state
	li.htmlEscape = previous.htmlEscape
	li.first = previous.first
	li.linebreak = previous.linebreak

	if li.first {
		processedStr = strings.Replace(processedStr, "&#124;", "|", -1)
		processedStr = strings.Replace(processedStr, "&#58;", ":", -1)
		processedStr = strings.Replace(processedStr, "&#40;", "(", -1)
		processedStr = strings.Replace(processedStr, "&#41;", ")", -1)
	}

	return strings.TrimSpace(processedStr)
}

type opencloseInterface interface {
	Pos() int
}

type opencloseSep int

func (op opencloseSep) Pos() int {
	return int(op)
}

type openclosersI []opencloseInterface

func (op openclosersI) Len() int {
	return len(op)
}

func (op openclosersI) Less(i, j int) bool {
	return op[i].Pos() < op[j].Pos()
}

func (op openclosersI) Swap(i, j int) {
	op[i], op[j] = op[j], op[i]
}

func (li *Lispy) parseParamExt(acontent string, sep rune) int {
	sep_indexes := indexRune(acontent, sep)

	openers_indexes := indexRune(acontent, openRune)
	closers_indexes := indexRune(acontent, closeRune)

	openers := openclosersI{}
	open := 0

	for _, sep := range sep_indexes {
		openers = append(openers, opencloseSep(sep))
	}

	for _, opener := range openers_indexes {
		openers = append(openers, openclose{true, opener})
	}

	for _, closer := range closers_indexes {
		openers = append(openers, openclose{false, closer})
	}

	sort.Sort(openers)

	pos := -1

	kill := false

	// Find seprater that outside of opener
	for _, opener := range openers {
		switch t := opener.(type) {
		case opencloseSep:
			if open <= 0 {
				pos = t.Pos()
				kill = true
			}
		case openclose:
			if t.opener {
				open++
			} else {
				open--
			}
		}
		if kill {
			break
		}
	}

	return pos
}

func (li *Lispy) ParseParam() {
	li.parseParam()
}

func (li *Lispy) parseParam() {
	if li.paramParsed {
		return
	}
	li.paramParsed = true
	li.Content = strings.TrimSpace(li.Content)

	pos := li.parseParamExt(li.Content, '|')

	if pos == -1 {
		return
	}

	params := li.Content[pos+1:]

	li.Content = strings.TrimSpace(li.Content[:pos])

	for {
		namepos := li.parseParamExt(params, ':')
		name := ""
		if namepos == -1 {
			name = params
		} else {
			name = params[:namepos]
		}
		params = params[namepos+1:]
		valuepos := li.parseParamExt(params, '|')
		value := ""
		if valuepos == -1 {
			if namepos == -1 {
				break
			}
			value = params
			li.Set(name, value)
			break
		} else {
			value = params[:valuepos]
		}
		params = params[valuepos+1:]
		li.Set(name, value)
	}

	li.filters()
}

func (li *Lispy) filters() {
	filters := li.Get("~filters")
	li.Delete("~filters")
	if filters == "" {
		return
	}
	if li.restrictParam {
		return
	}
	for _, filter := range strings.Split(filters, ",") {
		filter = strings.TrimSpace(filter)
		switch filter {
		case "line":
			li.Content = strings.Replace(li.Content, "\r\n", "(br:)", -1)
			li.Content = strings.Replace(li.Content, "\r", "(br:)", -1)
			li.Content = strings.Replace(li.Content, "\n", "(br:)", -1)
		case "tab":
			li.Content = strings.Replace(li.Content, "\t", "&nbsp;&nbsp;", -1)
		}
	}
}

// Set Parameter/Attribute
func (li *Lispy) Set(name, value string) {
	if li.param == nil {
		li.param = map[string][]string{}
	}

	name = strings.TrimSpace(strings.ToLower(name))

	name = strings.Replace(name, "&#58;", `:`, -1)

	if len(li.param[name]) <= 0 {
		li.param[name] = append(li.param[name], value)
	} else {
		li.param[name][0] = value
	}
}

// Get Parameter/Attribute
func (li *Lispy) Get(name string) string {
	if !li.Exist(name) {
		return ""
	}

	return li.param[name][0]
}

// Get Raw Parameter/Attribute (Unescaped)
func (li *Lispy) GetRaw(name string) string {
	return _html.UnescapeString(li.Get(name))
}

// Get Parameter/Attribute as Int64, return 0 on fail!
func (li *Lispy) GetInt64(name string) int64 {
	num, err := strconv.ParseInt(li.Get(name), 10, 64)
	if err != nil {
		num = 0
	}
	return num
}

// Get Parameter/Attribute as Int, return 0 on fail!
func (li *Lispy) GetInt(name string) int {
	return int(li.GetInt64(name))
}

// Check for existant of Parameter/Attribute
func (li *Lispy) Exist(name string) bool {
	li.parseParam()
	if len(li.param[name]) <= 0 {
		return false
	}

	return true
}

// Delete Parameter/Attribute
func (li *Lispy) Delete(name string) {
	if li.Exist(name) {
		delete(li.param, name)
	}
}

// Get Names of Parameter/Attribute.
//
// Note: Does not operate while Parameter/Attribute has been Restricted.
func (li *Lispy) GetNames() []string {
	li.parseParam()
	if li.restrictParam {
		return []string{}
	}

	str := []string{}
	for name, _ := range li.param {
		str = append(str, name)
	}

	sort.Sort(sort.StringSlice(str))

	return str
}

// HTML Render
func (li *Lispy) HtmlRender(htmlstr string) string {
	htmlfunc := html.FuncMap{
		"names": func() []string {
			return li.GetNames()
		},
		"get": func(name string) html.HTML {
			return li.Safe(li.Get(name))
		},
		"getint64": func(name string) int64 {
			return li.GetInt64(name)
		},
		"getint": func(name string) int {
			return li.GetInt(name)
		},
		"getdel": func(name string) html.HTML {
			defer li.Delete(name)
			return li.Safe(li.Get(name))
		},
		"getint64del": func(name string) int64 {
			defer li.Delete(name)
			return li.GetInt64(name)
		},
		"getintdel": func(name string) int {
			defer li.Delete(name)
			return li.GetInt(name)
		},
		"exist": func(name string) bool {
			return li.Exist(name)
		},
		"existdel": func(name string) bool {
			defer li.Delete(name)
			return li.Exist(name)
		},
		"existres": func(name string) bool {
			defer li.Delete(name)
			if li.restrictParam {
				return false
			}
			return li.Exist(name)
		},
		"render": func(str string) html.HTML {
			return li.RenderSafe(str)
		},
		"renderunsafe": func(str string) string {
			return li.Render(str)
		},
		"safe": func(str string) html.HTML {
			return li.Safe(str)
		},
		"css_safe": func(str string) html.CSS {
			return html.CSS(str)
		},
		"js_safe": func(str string) html.JS {
			return html.JS(str)
		},
		"attr": func(str string) html.HTMLAttr {
			return html.HTMLAttr(str)
		},
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()
	t, err := html.New("html").Funcs(htmlfunc).Parse(htmlstr)
	if err != nil {
		buf.WriteString(err.Error())
	}
	err = t.Execute(buf, li)
	if err != nil {
		buf.WriteString(err.Error())
	}
	return buf.String()
}

// Restrict Parameter, useful for comment system.
func (li *Lispy) RestrictParam() {
	li.restrictParam = true
}

// Restrict Attribute, useful for comment system.
//
// Note: Alias of RestrictParam
func (li *Lispy) RestrictAttribute() {
	li.RestrictParam()
}

// Specify Allowed Tags, useful for comment system.
//
// Note: All tags are enabled by default.
func (li *Lispy) AllowTags(names ...string) {
	li.allowedNames = append(li.allowedNames, names...)
}

// Get Content as Int64, return 0 on fail!
func (li *Lispy) ContentAsInt64() int64 {
	num, err := strconv.ParseInt(li.Content, 10, 64)
	if err != nil {
		num = 0
	}
	return num
}

// Disable Automatic HTML Escape
//
// Note: It's recommended that you don't disable that!
func (li *Lispy) DisableAutoEscape() {
	li.htmlEscape = false
}

// Enable Automatic Line Break, useful for comment system.
func (li *Lispy) EnableAutoLineBreak() {
	li.linebreak = true
}

// Unescaped Content
func (li *Lispy) RawContent() string {
	return _html.UnescapeString(li.Content)
}

func indexRune(str string, c rune) []int {
	indexs := []int{}
	for pos, cc := range str {
		if cc == c {
			indexs = append(indexs, pos)
		}
	}
	return indexs
}
