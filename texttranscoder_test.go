package wikiwiki_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-wikiwiki"
	"github.com/reiver/go-wikiwiki/renderer/html"
)

func TestTextTranscode(t *testing.T) {

	tests := []struct{
		Text string
		Expected string
	}{
		{
			Text:     "",
			Expected: "",
		},



		{
			Text:     "Hello world!",
			Expected: "Hello world!",
		},



		{
			Text:     "one ''two'' three",
			Expected: "one <sup>two</sup> three",
		},
		{
			Text:     "one ((two)) three",
			Expected: "one two three",
		},
		{
			Text:     "one **two** three",
			Expected: "one <b>two</b> three",
		},
		{
			Text:     "one ,,two,, three",
			Expected: "one <sub>two</sub> three",
		},
		{
			Text:     "one //two// three",
			Expected: "one <i>two</i> three",
		},
		{
			Text:     "one ::banana:: three",
			Expected: `one <span class="wiki-code">banana</span> three`,
		},
		{
			Text:     "one [[two]] three",
			Expected: `one <span class="wiki-link">two</span><script>(function(){`+
				`e=document.querySelector(".wiki-link");`+
				`n=document.createElement("a");`+
				`t=e.innerText;`+
				`n.href=t;`+
				`n.innerText=t;`+
				`e.replaceWith(n)`+
			`})()</script> three`,
		},
		{
			Text:     "one __two__ three",
			Expected: "one <u>two</u> three",
		},
		{
			Text:     "one ``two`` three",
			Expected: "one <tt>two</tt> three",
		},
		{
			Text:     "one {{two}} three",
			Expected: `one <span class="wiki-name-value">two</span><script>(function(){`+
				`e=document.querySelector(".wiki-name-value");`+
				`m=document.getElementById("local."+e.innerText);`+
				`if(m){`+
					`t=m.getAttribute("content");`+
					`n=document.createElement("span");`+
					`n.innerText=t;`+
					`e.replaceWith(n)`+
				`}`+
			`})()</script> three`,
		},
		{
			Text:     "one ||two|| three",
			Expected: "one <mark>two</mark> three",
		},
		{
			Text:     "one ‾‾two‾‾ three",
			Expected: `one <span style="text-decoration:overline">two</span> three`,
		},



		{
			Text:     "e''iπ''",
			Expected: "e<sup>iπ</sup>",
		},



		{
			Text:     "``apple **b||ana||na** //cherry//``",
			Expected: "<tt>apple <b>b<mark>ana</mark>na</b> <i>cherry</i></tt>",
		},



		{
			Text:     "2''3((''4''))''",
			Expected: "2<sup>3<sup>4</sup></sup>",
		},



		{
			Text:     "apple ''ban'ana'' cherry",
			Expected: "apple <sup>ban'ana</sup> cherry",
		},
	}

	testloop: for testNumber, test := range tests {

		var buffer strings.Builder
		renderer := wikiwikihtml.NewRenderer()

		var texttranscoder wikiwiki.TextTranscoder = wikiwiki.NewTextTranscoder(&buffer, renderer)

		for runeNumber, r := range test.Text {
			err := texttranscoder.InterpretRune(r)
			if nil != err {
				t.Errorf("For test #%d and rune #%d, did not expect an error when sending rune but actually got one." , testNumber, runeNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("RUNE: %q (%U)", r, r)
				t.Logf("TEXT: %q", test.Text)
				continue testloop
			}
		}
		{
			err := texttranscoder.Close()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error when closing but actually got one." , testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("TEXT: %q", test.Text)
				continue testloop
			}
		}

		{
			expected := test.Expected
			actual   := buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual rendered result is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("TEXT:     %q", test.Text)
				continue testloop
			}
		}
	}
}
