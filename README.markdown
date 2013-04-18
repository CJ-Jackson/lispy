# Lispy Syntax System

Simple, Elegant yet Extendible Syntax System.

## Syntax ##

	(tagname`content|attr1:value1|attr2:value2|etc:etc`)

### Escapers ###

	\| and  \:

Note: Brackets '()' and Backticks '`' do not need escaping, just don't match the syntax!

## Installation ##

	go get github.com/CJ-Jackson/lispy
	
## Example ##

	(video`(source`forrest_gump.mp4|type:video/mp4`)
		(source`forrest_gump.ogg|type:video/ogg`)
		(track`subtitles_en.vtt|kind:subtitles|srclang:en|label:English`)
		(track`subtitles_no.vtt|kind:subtitles|srclang:no|label:Norwegian`)
	|width:320|height:240|controls:`)

Output:

	<video width="320" height="240" controls><source src="forrest_gump.mp4" type="video/mp4" />
		<source src="forrest_gump.ogg" type="video/ogg" />
		<track kind="subtitles" label="English" src="subtitles_en.vtt" srclang="en" />
		<track kind="subtitles" label="Norwegian" src="subtitles_no.vtt" srclang="no" />
	</video>

## Usage ##

	package main

	import (
		"github.com/CJ-Jackson/lispy"
		"fmt"
	)

	func main() {
		parser := lispy.New()
		fmt.Print(parser.Parse("(p`Hello World!`)"))
		// Output: <p>Hello World!</p>
	}

## Documentation ##

Documentation are avaliable at

http://go.pkgdoc.org/github.com/CJ-Jackson/lispy

## Note ##

Like HTML it's is suspectible to leakage when forgetting to close tag, unlike HTML it is sandboxed to a variable, while HTML is sandboxed to a page.  