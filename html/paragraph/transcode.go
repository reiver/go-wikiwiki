package paragraph

import (
	"errors"
	"io"

	"github.com/reiver/go-unicode"
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-utf8"

	"github.com/reiver/go-wikiwiki/html/renderer/text"
	"github.com/reiver/go-wikiwiki/renderer"
	"github.com/reiver/go-wikiwiki/transcoder"
	"github.com/reiver/go-wikiwiki/transcoder/text"
)

func Transcode(writer io.Writer, reader io.Reader) (err error) {
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


        var textrenderer wikiwikirenderer.TextRenderer = wikiwikihtmltextrenderer.TextRenderer
        if nil == textrenderer {
                return errInternalError
        }

        var texttranscoder wikiwikitranscoder.TextTranscoder = wikiwikitexttranscoder.NewTextTranscoder(writer, textrenderer)
        if nil == texttranscoder {
                return errInternalError
        }
        defer func() {
                err = texttranscoder.Close()
                if nil != err {
                        err = erorr.Errorf("wikiwiki: problem closing text-transcoder: %w", err)
                }
        }()

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
			default:
				err = texttranscoder.InterpretRune(r)
				if nil != err {
					return erorr.Errorf("wikiwiki: text-transcoder had trouble interpretting rune %q (%U): %w", r, r, err)
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
