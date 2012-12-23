package lispy

import (
	"fmt"
	"testing"
)

func TestMedia(t *testing.T) {
	fmt.Println("Media Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(video`(source`forrest_gump.mp4|type:video/mp4`)\r\n"
	code += "(source`forrest_gump.ogg|type:video/ogg`)\r\n"
	code += "(track`subtitles_en.vtt|kind:subtitles|srclang:en|label:English`)\r\n"
	code += "(track`subtitles_no.vtt|kind:subtitles|srclang:no|label:Norwegian`)\r\n"
	code += "|width:320|height:240|controls:`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(audio`(source`horse.ogg|type:audio/ogg`)\r\n"
	code += "(source`horse.mp3|type:audio/mpeg`)\r\n"
	code += "(track`subtitles_en.vtt|kind:subtitles|srclang:en|label:English`)\r\n"
	code += "(track`subtitles_no.vtt|kind:subtitles|srclang:no|label:Norwegian`)\r\n"
	code += "|width:320|height:240|controls:`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
