package wikiwikihtml_test

import (
	"testing"

	"io"
	"io/fs"
	"time"

	"sourcecode.social/reiver/go-strfs"

	"github.com/reiver/go-wikiwiki/html"
)

func TestFile(t *testing.T) {

	tests := []struct{
		Content string
		Expected string
	}{
		{
			Content:  "",
			Expected: "",
		},



		{
			Content:
				"Hello **world**!"+"\n"+
				"**How** are //you//?"+"\n"+
				"How ||are you|| today?"+"\n",
			Expected:
				"Hello <b>world</b>!"+"\n"+
				"<b>How</b> are <i>you</i>?"+"\n"+
				"How <mark>are you</mark> today?"+"\n",
		},
	}

	for testNumber, test := range tests {

		var content strfs.Content = strfs.CreateContent(test.Content)
		var regularfile strfs.RegularFile = strfs.RegularFile{
			FileContent: content,
			FileName:    "helloworld.wiki",
			FileModTime: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
		}
		var wikifile fs.File = &regularfile

		var htmlfile fs.File = wikiwikihtml.File(wikifile)

		actualBytes, err := io.ReadAll(htmlfile)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("CONTENT: %q", test.Content)
			continue
		}

		err = htmlfile.Close()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("CONTENT: %q", test.Content)
			continue
		}

		{
			var actual   string = string(actualBytes)
			var expected string = test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual 'html content' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
		}
	}
}
