package wikiwikihtml_test

import (
	"fmt"
	"io"
	"strings"

	"github.com/reiver/go-wikiwiki/html"
)

func ExampleTranscode() {

	var wikitext string =
		"Hello world!"+"\n"+
		"This is the **second** sentence."+"\n"+
		"Did you know that: 0 = 1 + e''iπ''."+"\n"+
		"This is the //fourth// sentence."+"\n"+
		"This is the ``fifth`` sentence."+"\n"+
		"Apple ||BANANA|| cherry."+"\n"+
		"This is the ||7''th''|| sentence."+"\n"+
		"‾‾x‾‾ = (x,,1,,, x,,2,,, x,,3,,)"+"\n"

	var input io.Reader = strings.NewReader(wikitext)

	var output strings.Builder

	err := wikiwikihtml.Transcode(&output, input)
	if nil != err {
		panic(err)
	}

	fmt.Print(output.String())

	// Output:
	//
	// Hello world!
	// This is the <b>second</b> sentence.
	// Did you know that: 0 = 1 + e<sup>iπ</sup>.
	// This is the <i>fourth</i> sentence.
	// This is the <tt>fifth</tt> sentence.
	// Apple <mark>BANANA</mark> cherry.
	// This is the <mark>7<sup>th</sup></mark> sentence.
	// <span style="text-decoration:overline">x</span> = (x<sub>1</sub>, x<sub>2</sub>, x<sub>3</sub>)
}
