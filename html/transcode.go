package wikiwikihtml

import (
	"errors"
	"io"

	"github.com/reiver/go-iopsep"
	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-wikiwiki/html/renderer/text"
	"github.com/reiver/go-wikiwiki/internal/magic"
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
		_, err := wikiwikimagic.ReadMagic(reader)
		if nil != err {
			return err
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
		var block io.ReadCloser = iopsep.NewParagraphReadCloser(reader)
		if nil == block {
			return errNilBlockReader
		}

		var r rune

		// ignore any leading empty lines.
		loop: for {
			var size int

			r, size, err = utf8.ReadRune(block)
			if 0 < size {
				switch r {
				case unicode.LF, unicode.CR, unicode.NEL, unicode.PS:
					// nothing here.
				default:
					break loop
				}
			}
			if io.EOF == err {
				return nil
			}
			if nil != err {
				return erorr.Errorf("wikiwiki: problem reading rune: %w", err)
			}
		}

		var  element internalElement = internalElement(string(r))

		{
			switch element {
			case "§":
				miniloop: for {
					var size int

					r, size, err = utf8.ReadRune(block)
					if 0 < size {
						switch r {
						case '§':
							element = element+"§"
							continue
						case ' ':
							break miniloop
						default:
							return erorr.Error("wikiwiki: heading requires a space between § symbols and title text")
						}
					}

				}
			}
		}

		{
			var code string = element.Begin()

			_, err := io.WriteString(writer, code)
			if nil != err {
				return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
			}
		}

		{
			var buffer string = element.Buffer()
			if 0 < len(buffer) {
				_, err := io.WriteString(writer, buffer)
				if nil != err {
					return erorr.Errorf("wikiwiki: problem writing %q: %w", buffer, err)
				}
			}
		}

		for {

			r, size, err := utf8.ReadRune(block)
			if 0 < size {

				var particle internalParticle = internalParticle(r)

				render, rendered := particle.Render()
				if rendered {
					_, err := io.WriteString(writer, render)
					if nil != err {
						return erorr.Errorf("wikiwiki: problem writing %q: %w", render, err)
					}
				} else {
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
				return erorr.Errorf("wikiwiki: problem readung rune: %w", err)
			}
		}

		{
			var code string = element.End()

			_, err := io.WriteString(writer, code)
			if nil != err {
				return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
			}
			element = ""
		}
	}

	return nil
}
