package lispy

var Map = LispyMap{
	"--":         Blanks,
	"a":          A,
	"abbr":       CommonTitle,
	"address":    Common,
	"area":       Area,
	"article":    Common,
	"aside":      Common,
	"audio":      Audio,
	"b":          Common,
	"bdi":        Common,
	"bdo":        CommonDir,
	"blockquote": CommonCite,
	"br":         Br,
	"button":     Button,
	"canvas":     Canvas,
	"caption":    Common,
	"cite":       Common,
	"code":       Common,
	"col":        CommonSpan,
	"colgroup":   CommonSpan,
	"css":        CSS,
	"datalist":   Common,
	"dd":         Common,
	"del":        CommonCiteDateTime,
	"details":    Details,
	"dfn":        Common,
	"div":        Common,
	"dl":         Common,
	"dt":         Common,
	"em":         Common,
	"embed":      Embed,
	"fieldset":   FieldSet,
	"figcaption": Common,
	"figure":     Common,
	"footer":     Common,
	"form":       Form,
	"h1":         Header,
	"h2":         Header,
	"h3":         Header,
	"h4":         Header,
	"h5":         Header,
	"h6":         Header,
	"header":     Common,
	"hgroup":     Common,
	"hr":         Common,
	"i":          Common,
	"iframe":     CommonSrc,
	"img":        Img,
	"input":      Input,
	"ins":        CommonCiteDateTime,
	"javascript": JavaScript,
	"kbd":        Common,
	"keygen":     Keygen,
	"label":      Label,
	"legend":     Common,
	"li":         CommonValue,
	"map":        CommonName,
	"mark":       Common,
	"nav":        Common,
	"object":     Object,
	"ol":         Ol,
	"option":     Option,
	"output":     Output,
	"p":          Common,
	"param":      Param,
	"pre":        Common,
	"q":          CommonCite,
	"rp":         Common,
	"rt":         Common,
	"ruby":       Common,
	"s":          Common,
	"samp":       Common,
	"section":    Common,
	"select":     Select,
	"small":      Common,
	"source":     Source,
	"span":       Common,
	"strong":     Common,
	"sub":        Common,
	"sup":        Common,
	"table":      Table,
	"tbody":      Common,
	"td":         TableTd,
	"textarea":   Textarea,
	"tfood":      Common,
	"th":         TableTd,
	"thead":      Common,
	"time":       Time,
	"tr":         Common,
	"track":      Track,
	"u":          Common,
	"ul":         Common,
	"var":        Common,
	"video":      Video,
	"wbr":        Wbr,
}

var HandlerMap = LispyHandlerMap{}
