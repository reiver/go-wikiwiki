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
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				""+
				"</body>\n</html>\n",
		},



		{
			Content:  "&",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"&amp;"+
				"</p>\n"+
				"</body>\n</html>\n",
		},
		{
			Content:  "<",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"&lt;"+
				"</p>\n"+
				"</body>\n</html>\n",
		},
		{
			Content:  ">",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"&gt;"+
				"</p>\n"+
				"</body>\n</html>\n",
		},



		{
			Content:  "2 < 5 & 4 > 2 & 13 > 7",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"2 &lt; 5 &amp; 4 &gt; 2 &amp; 13 &gt; 7"+
				"</p>\n"+
				"</body>\n</html>\n",
		},



		{
			Content:
				"Hello **world**!"+"\n"+
				"**How** are //you//?"+"\n"+
				"How ||are you|| today?"+"\n",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"Hello <b>world</b>!"+"\n"+
				"<b>How</b> are <i>you</i>?"+"\n"+
				"How <mark>are you</mark> today?"+"\n"+
				"</p>\n"+
				"</body>\n</html>\n",
		},



		{
			Content:
				"Hello **world**!"+"\n"+
				"**How** are //you//?"+"\n"+
				"How ||are you|| today?"+"\n"+
				""+"\n"+
				"apple banana cherry!"+"\n",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"Hello <b>world</b>!"+"\n"+
				"<b>How</b> are <i>you</i>?"+"\n"+
				"How <mark>are you</mark> today?"+"\n"+
				""+"\n"+
				"</p>\n"+
				"<p>\n"+
				"apple banana cherry!"+"\n"+
				"</p>\n"+
				"</body>\n</html>\n",
		},



		{
			Content:
				"Hello **world**!"+"\n"+
				"**How** are //you//?"+"\n"+
				"How ||are you|| today?"+"\n"+
				""+"\n"+
				"apple banana cherry!"+"\n"+
				""+"\n"+
				"― Do your ears hang low?"+"\n"+
				"― Do they wobble to and fro?"+"\n"+
				"― Can you tie 'em in a knot?"+"\n"+
				"― Can you tie 'em in a bow?"+"\n",
			Expected:
				"<html>\n<head>\n<meta charset=\"utf-8\" />\n<style type=\"text/css\">\na.wiki-link::before{\ncontent:attr(href);\n}\n</style>\n</head>\n<body>\n"+
				"<p>\n"+
				"Hello <b>world</b>!"+"\n"+
				"<b>How</b> are <i>you</i>?"+"\n"+
				"How <mark>are you</mark> today?"+"\n"+
				""+"\n"+
				"</p>\n"+
				"<p>\n"+
				"apple banana cherry!"+"\n"+
				""+"\n"+
				"</p>\n"+
				"<blockquote>\n"+
				"― Do your ears hang low?"+"\n"+
				"― Do they wobble to and fro?"+"\n"+
				"― Can you tie 'em in a knot?"+"\n"+
				"― Can you tie 'em in a bow?"+"\n"+
				"</blockquote>\n"+
				"</body>\n</html>\n",
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
