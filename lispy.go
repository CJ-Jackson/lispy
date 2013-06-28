// Simple, Elegant yet Extendible Syntax System. (p:Hello World in Paragraph!)
package lispy

import (
	"bytes"
	"fmt"
	_html "html"
	html "html/template"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// (name:value), (name:(name:value)), (name:content|key:value|key:value)

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
	CacheName     string
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
		CacheName:     "",
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

	newli := &Lispy{}
	*newli = *li
	newli.RWMutex = sync.RWMutex{}
	newli.param = map[string][]string{}
	newli.CacheName = ""

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

	searchSyntax := func() *syntaxPos {
		for key, value := range str {
			if value != ':' || key == 0 {
				continue
			}
			openPos := strings.LastIndex(str[:key], "(")
			if openPos == -1 {
				continue
			}
			if strings.ContainsRune(str[openPos+1:key], ' ') {
				continue
			}
			value := strings.ToLower(str[openPos+1 : key])
			if len(value) <= 0 {
				continue
			}
			return &syntaxPos{openPos, key, value}
		}

		return nil
	}

	for {
		index := searchSyntax()
		if index == nil {
			processedStr += str
			break
		}

		pos := index.Open

		lenght := index.Colon + 1
		li.Name = index.Name

		content_pos := lenght
		li.Content = str[content_pos:]

		content_lenght := len(li.Content)
		open := 1

		// Find closer for current opener!
		for key, value := range li.Content {
			if value == openRune {
				open++
			} else if value == closeRune {
				open--
			}
			if open <= 0 {
				content_lenght = key
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
		li.CacheName = ""

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
		processedStr = unescape(processedStr)
	}

	return strings.TrimSpace(processedStr)
}

func (li *Lispy) RenderedContent() html.HTML {
	return html.HTML(li.Render(li.Content))
}

func unescape(str string) string {
	str = strings.Replace(str, "&#124;", "|", -1)
	str = strings.Replace(str, "&#58;", ":", -1)
	str = strings.Replace(str, "&#40;", "(", -1)
	str = strings.Replace(str, "&#41;", ")", -1)
	return str
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

	poss := stringPoss{}

	open := 0

	barFound := false

	for key, value := range li.Content {
		switch value {
		case '|':
			if open <= 0 {
				poss = append(poss, stringPos{'|', key})
				barFound = true
			}
		case ':':
			if open <= 0 && barFound {
				poss = append(poss, stringPos{':', key})
			}
		case openRune:
			open++
		case closeRune:
			open--
		}
	}

	if len(poss) == 0 {
		return
	}

	for key, sep := range poss {
		if sep.Char != '|' {
			continue
		}
		colonPos := -1
		barPos := -1
		for _, _sep := range poss[key+1:] {
			if _sep.Char == ':' && colonPos == -1 {
				colonPos = _sep.Pos
			} else if _sep.Char == '|' && colonPos != -1 {
				barPos = _sep.Pos
			}

			if colonPos != -1 && barPos != -1 {
				break
			}
		}
		var value string
		if barPos != -1 {
			value = li.Content[colonPos+1 : barPos]
		} else {
			value = li.Content[colonPos+1:]
		}
		if colonPos == -1 {
			li.Set(li.Content[sep.Pos+1:], "")
		} else {
			li.Set(li.Content[sep.Pos+1:colonPos], value)
		}

	}

	li.Content = li.Content[:poss[0].Pos]

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
	li.parseParam()
	if !li.Exist(name) {
		return ""
	}

	return li.param[name][0]
}

func (li *Lispy) GetDel(name string) string {
	defer li.Delete(name)
	return li.Get(name)
}

// Get Raw Parameter/Attribute (Unescaped)
func (li *Lispy) GetRaw(name string) string {
	return _html.UnescapeString(li.Get(name))
}

func (li *Lispy) GetRawDel(name string) string {
	defer li.Delete(name)
	return li.GetRaw(name)
}

// Get Parameter/Attribute as Int64, return 0 on fail!
func (li *Lispy) GetInt64(name string) int64 {
	num, err := strconv.ParseInt(li.Get(name), 10, 64)
	if err != nil {
		num = 0
	}
	return num
}

func (li *Lispy) GetInt64Del(name string) int64 {
	defer li.Delete(name)
	return li.GetInt64(name)
}

// Get Parameter/Attribute as Int, return 0 on fail!
func (li *Lispy) GetInt(name string) int {
	return int(li.GetInt64(name))
}

func (li *Lispy) GetIntDel(name string) int {
	defer li.Delete(name)
	return li.GetInt(name)
}

// Check for existant of Parameter/Attribute
func (li *Lispy) Exist(name string) bool {
	li.parseParam()
	if len(li.param[name]) <= 0 {
		return false
	}

	return true
}

func (li *Lispy) ExistDel(name string) bool {
	defer li.Delete(name)
	return li.Exist(name)
}

func (li *Lispy) ExistRes(name string) bool {
	defer li.Delete(name)
	if li.restrictParam {
		return false
	}
	return li.Exist(name)
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

func (li *Lispy) GetParam() string {
	str := ""

	for _, name := range li.GetNames() {
		str += fmt.Sprintf(` %s="%s"`, name, li.Get(name))
	}

	return str
}

var _templateCache = struct {
	sync.RWMutex
	c map[string]*html.Template
}{
	c: map[string]*html.Template{},
}

// HTML Render
func (li *Lispy) HtmlRender(htmlstr string) string {
	if li.CacheName == "" {
		li.CacheName = li.Name
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	var err error

	_templateCache.RLock()
	t := _templateCache.c[li.CacheName]
	_templateCache.RUnlock()
	if t == nil {
		t, err = html.New("html").Funcs(html.FuncMap{
			"html": func(str string) html.HTML {
				return html.HTML(str)
			},
			"css": func(str string) html.CSS {
				return html.CSS(str)
			},
			"js": func(str string) html.JS {
				return html.JS(str)
			},
			"attr": func(str string) html.HTMLAttr {
				return html.HTMLAttr(str)
			},
		}).Parse(htmlstr)
		if err != nil {
			buf.WriteString(err.Error())
			return buf.String()
		}
		_templateCache.Lock()
		_templateCache.c[li.CacheName] = t
		_templateCache.Unlock()
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

// Unescaped Content Extended, includes | : ( )
func (li *Lispy) RawContentExt() string {
	return unescape(li.RawContent())
}

type syntaxPos struct {
	Open, Colon int
	Name        string
}

type stringPos struct {
	Char rune
	Pos  int
}

type stringPoss []stringPos

func (str stringPoss) Len() int {
	return len(str)
}

func (str stringPoss) Less(i, j int) bool {
	return str[i].Pos < str[j].Pos
}

func (str stringPoss) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
