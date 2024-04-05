package paragraph

import (
	"errors"
	"io"

	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"
)

func Transcode(writer io.Writer, reader io.Reader) error {
	if nil == writer {
		return errNilWriter
	}
	if nil == reader {
		return errNilReader
	}

	{
		const code string = "<p>"

		_, err := io.WriteString(writer, code)
		if nil != err {
			return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing %q: %w", code, err)
		}
	}

	for {
		r, size, err := utf8.ReadRune(reader)
		if 0 < size {

			switch r {
			case unicode.LS:
				const code string = "<br />"

				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing %q: %w", code, err)
				}
			case '&':
				const code string = "&amp;"

				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing %q: %w", code, err)
				}
			case '<':
				const code string = "&lt;"

				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing %q: %w", code, err)
				}
			case '>':
				const code string = "&gt;"

				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing %q: %w", code, err)
				}
			default:
				_, err := utf8.WriteRune(writer, r)
				if nil != err {
					return erorr.Errorf("wikiwiki: while trying to transcode wiki to HTML, had problem writing rune %q (%U): %w", r, r, err)
				}
			}

		}
		if errors.Is(err, io.EOF) {
	/////////////// BREAK
			break
		}
		if nil != err {
			return erorr.Errorf("wikiwiki: trying to transcode wiki to HTML, had problem reading a rune: %w", err)
		}
	}

	{
		const code string = "</p>\n"

		_, err := io.WriteString(writer, code)
		if nil != err {
			erorr.Errorf("wikiwiki: problem %q: %w", code, err)
		}
	}

	return nil
}
