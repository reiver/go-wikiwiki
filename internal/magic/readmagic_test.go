package wikiwikimagic

import (
	"testing"

	"io"
	"strings"

	. "github.com/reiver/go-wikiwiki/internal/eol"
)

func TestReadMagic(t *testing.T) {

	tests := []struct{
		Data string
		Expected int
	}{
		{
			Data:         "wiki/1",
			Expected: len("wiki/1"),
		},



		{
			Data:         "wiki/1"  +     string(LF),
			Expected: len("wiki/1") + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(VT),
			Expected: len("wiki/1") + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(FF),
			Expected: len("wiki/1") + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(CR),
			Expected: len("wiki/1") + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(NEL),
			Expected: len("wiki/1") + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(LS),
			Expected: len("wiki/1") + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(PS),
			Expected: len("wiki/1") + len(string(PS)),
		},



		{
			Data:         "wiki/1"  +     string(LF)+string(CR),
			Expected: len("wiki/1") + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(CR)+string(LF),
			Expected: len("wiki/1") + len(string(CR)),
		},



		{
			Data:         "wiki/1"  +     string(LF)+"Hello world!",
			Expected: len("wiki/1") + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(VT)+"Hello world!",
			Expected: len("wiki/1") + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(FF)+"Hello world!",
			Expected: len("wiki/1") + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(CR)+"Hello world!",
			Expected: len("wiki/1") + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(NEL)+"Hello world!",
			Expected: len("wiki/1") + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(LS)+"Hello world!",
			Expected: len("wiki/1") + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(PS)+"Hello world!",
			Expected: len("wiki/1") + len(string(PS)),
		},



		{
			Data:         "wiki/1"  +     string(HT),
			Expected: len("wiki/1") + len(string(HT)),
		},
		{
			Data:         "wiki/1"  +     string(SP),
			Expected: len("wiki/1") + len(string(SP)),
		},

		{
			Data:         "wiki/1"  +     string(NBSP),
			Expected: len("wiki/1") + len(string(NBSP)),
		},



		{
			Data:         "wiki/1"  +     string(HT)  +     string(LF)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(VT)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(FF)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(CR)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(NEL)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(LS)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(HT)  +     string(PS)+"Hello world!",
			Expected: len("wiki/1") + len(string(HT)) + len(string(PS)),
		},



		{
			Data:         "wiki/1"  +     string(SP)  +     string(LF)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(VT)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(FF)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(CR)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(NEL)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(LS)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     string(PS)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len(string(PS)),
		},



		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(LF)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(VT)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(FF)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(CR)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(NEL)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(LS)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(NBSP)  +     string(PS)+"Hello world!",
			Expected: len("wiki/1") + len(string(NBSP)) + len(string(PS)),
		},



		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(LF)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(LF)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(VT)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(VT)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(FF)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(FF)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(CR)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(CR)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(NEL)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(NEL)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(LS)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(LS)),
		},
		{
			Data:         "wiki/1"  +     string(SP)  +     "one = 1"  +     string(PS)+"Hello world!",
			Expected: len("wiki/1") + len(string(SP)) + len("one = 1") + len(string(PS)),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)

		actual, err := ReadMagic(reader)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %q", test.Data)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes-read is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
		}
	}
}

func TestReadMagic_fail(t *testing.T) {

	tests := []struct{
		Data string
	}{
		{
			Data: "WIKI/1",
		},



		{
			Data: "wiki/2",
		},



		{
			Data: "banana/3",
		},



		{
			Data: "Hello world!",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)

		actual, err := ReadMagic(reader)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("DATA: %q", test.Data)
			continue
		}

		{
			const expected int = 6

			if expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes-read is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
		}
	}
}
