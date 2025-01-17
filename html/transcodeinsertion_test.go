package wikiwikihtml

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestTranscodeInsertion(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value: "\r⁁ img.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "\r⁁\timg.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},



		{
			Value: "\r⁁ vid.mp4",
			Expected: `<figure><video src="vid.mp4" controls></video></figure>`+"\n",
		},
		{
			Value: "\r⁁\tvid.mp4",
			Expected: `<figure><video src="vid.mp4" controls></video></figure>`+"\n",
		},



		{
			Value: "\r⁁ doc.txt",
			Expected: `<figure><a href="doc.txt" />doc.txt</a></figure>`+"\n",
		},
		{
			Value: "\r⁁\tdoc.txt",
			Expected: `<figure><a href="doc.txt" />doc.txt</a></figure>`+"\n",
		},



		{
			Value: "⁁ img.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁\timg.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},



		{
			Value: "⁁  img.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ \timg.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁\t img.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁\t\timg.png",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},



		{
			Value: "⁁ img.png"+"\t",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+"\n",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+"\r",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+" ",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+"\u0085",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+"\u2028",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
		{
			Value: "⁁ img.png"+"\u2029",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},



		{
			Value: "⁁ img.png"+"\n\n"+"apple banana cherry",
			Expected: `<figure><img src="img.png" /></figure>`+"\n",
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		var runescanner io.RuneScanner = utf8.NewRuneScanner(strings.NewReader(test.Value))

		err := transcodeInsertion(&actualBuffer, runescanner)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual transcoded value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
	t.Log()
	t.Logf("Num-Tests: %d", len(tests))
}
