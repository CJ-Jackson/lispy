package lispy

import (
	"testing"
)

func TestMedia(t *testing.T) {
	lisp := New()
	code := "(video:(source:forrest_gump.mp4|type:video/mp4)"
	code += "(source:forrest_gump.ogg|type:video/ogg)"
	code += "(track:subtitles_en.vtt|kind:subtitles|srclang:en|label:English)"
	code += "(track:subtitles_no.vtt|kind:subtitles|srclang:no|label:Norwegian)"
	code += "|width:320|height:240|controls:)"
	str := lisp.Render(code)

	if str != `<video width="320" height="240" controls><source src="forrest_gump.mp4" type="video/mp4"/><source src="forrest_gump.ogg" type="video/ogg"/><track kind="subtitles" label="English" src="subtitles_en.vtt" srclang="en"/><track kind="subtitles" label="Norwegian" src="subtitles_no.vtt" srclang="no"/></video>` {
		t.Fail()
	}

	code = "(audio:(source:horse.ogg|type:audio/ogg)"
	code += "(source:horse.mp3|type:audio/mpeg)"
	code += "(track:subtitles_en.vtt|kind:subtitles|srclang:en|label:English)"
	code += "(track:subtitles_no.vtt|kind:subtitles|srclang:no|label:Norwegian)"
	code += "|controls:)"
	str = lisp.Render(code)

	if str != `<audio controls><source src="horse.ogg" type="audio/ogg"/><source src="horse.mp3" type="audio/mpeg"/><track kind="subtitles" label="English" src="subtitles_en.vtt" srclang="en"/><track kind="subtitles" label="Norwegian" src="subtitles_no.vtt" srclang="no"/></audio>` {
		t.Fail()
	}
}
