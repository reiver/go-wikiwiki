package wikiwiki_test

import (
	"fmt"
	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-wikiwiki"
	"github.com/reiver/go-wikiwiki/renderer/html"
)

func ExampleTranscodeBlock() {

	var wikitext string =
		"Hell,,o,, w''or''ld!"+"\n"+
		"This is the **second** sentence."+"\n"+
		"Did you know that: 0 = 1 + e''iπ''."+"\n"+
		""+"\n"+
		"This is the //fourth// sentence."+"\n"+
		"This is the ``fifth`` sentence."+"\n"+
		"Apple ||BANANA|| cherry."+"\n"+
		""+"\n"+
		"This is the ||7''th''|| sentence."+"\n"+
		"‾‾x‾‾ = (x,,1,,, x,,2,,, x,,3,,)"+"\n"

	var reader io.Reader = strings.NewReader(wikitext)
	var input io.RuneScanner = utf8.NewRuneScanner(reader)

	var output strings.Builder

	var renderer wikiwiki.Renderer = wikiwikihtml.NewRenderer()

	err := wikiwiki.TranscodeBlock(&output, renderer, input)
	if nil != err {
		panic(err)
	}

	fmt.Print(output.String())

	// Output:
	//
	// Hell<sub>o</sub> w<sup>or</sup>ld!
	// This is the <b>second</b> sentence.
	// Did you know that: 0 = 1 + e<sup>iπ</sup>.
}
