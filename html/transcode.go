package wikiwikihtml

import (
	"errors"
	"io"

	"github.com/reiver/go-iopsep"
	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"

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

	var closecode string

	for {
		var block io.ReadCloser = iopsep.NewParagraphReadCloser(reader)
		if nil == block {
			return errNilBlockReader
		}

		var firstrune bool = true

		var interpretted int


		for {
			r, size, err := utf8.ReadRune(block)
			if 0 < size {
				if firstrune {
					var opencode string

					switch r {
					case '§': // U+00A7 Section Sign
						opencode  = "<h2>"
						closecode = "</h2>\n"
					case '―': // U+2015 Horizontal Bar; i.e., quotation dash.
						opencode  = "<blockquote>\n"
						closecode = "</blockquote>\n"
					case '•', // U+2022 Bullet
					     '‣', // U+2023 Triangular Bullet
					     '⁃': // U+2043 Hyphen Bullet
						opencode  = "<ul>\n"
						closecode = "</ul>\n"
					default:
						opencode  = "<p>\n"
						closecode = "</p>\n"
					}


					firstrune = false

					_, err := io.WriteString(writer, opencode)
					if nil != err {
						return erorr.Errorf("wikiwiki: problem writing %q: %w", opencode, err)
					}
				}
			}
			if errors.Is(err, io.EOF) {
		/////////////// BREAK
				break
			}
			if nil != err {
				return erorr.Errorf("wikiwiki: problem readung rune: %w", err)
			}

			switch {
			case unicode.LS == r:
				const code string = "<br />\n"
				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
				}
			case "</ul>\n" == closecode && (
			         r == '•' || // U+2022 Bullet
			         r == '‣' || // U+2023 Triangular Bullet
			         r == '⁃' ): // U+2043 Hyphen Bullet
				const code string = "<li>"
				_, err := io.WriteString(writer, code)
				if nil != err {
					return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
				}
			default:
				err = texttranscoder.InterpretRune(r)
				if nil != err {
					return erorr.Errorf("wikiwiki: text-transcoder had trouble interpretting rune %q (%U): %w", r, r, err)
				}
			}
			interpretted++
		}

		if interpretted <= 0 {
			return nil
		}

		{
			_, err := io.WriteString(writer, closecode)
			if nil != err {
				return erorr.Errorf("wikiwiki: problem writing %q: %w", closecode, err)
			}
			closecode = ""
		}
	}

	return nil
}
