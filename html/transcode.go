package wikiwikihtml

import (
	"bufio"
	"errors"
	"io"

	"github.com/reiver/go-iolsep"
	"github.com/reiver/go-iopsep"
	"github.com/reiver/go-peek"
	"github.com/reiver/go-skip"
	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-wikiwiki/html/renderer/text"
	wikiwikifilemagic "github.com/reiver/go-wikiwiki/internal/magic"
	"github.com/reiver/go-wikiwiki/magic"
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
		_, err := wikiwikifilemagic.ReadMagic(reader)
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

	var bufferedreader *bufio.Reader = bufio.NewReader(reader)
	if nil == bufferedreader {
		return errNilBufferedReader
	}


	for {
		{
			err := skip.SkipRunes(bufferedreader, unicode.LF, unicode.CR, unicode.NEL, unicode.PS)
			if io.EOF == err {
				return nil
			}
			if nil != err {
				return erorr.Errorf("wikiwiki: problem skipping EOL runes: %w", err)
			}
		}

		var peeked rune
		{
			r, size, err := peek.PeekRune(bufferedreader)

			if nil != err {
				if 0 < size {
					err = texttranscoder.InterpretRune(r)
					if nil != err {
						return erorr.Errorf("wikiwiki: text-transcoder had trouble interpretting rune %q (%U): %w", r, r, err)
					}
				}
				if io.EOF == err {
					return nil
				}
				return err
			}

			if size <= 0 {
				return erorr.Errorf("wikiwiki: could not read rune at the beginning of the block: %w", err)
			}

			peeked = r
		}

		var element internalElement = internalElement(string(peeked))

		{

			const magicStringHeading string = string(wikiwikimagic.Heading)

			const magicStringRoundBulletedList      string = string(wikiwikimagic.RoundBulletedList)
			const magicStringTriangularBulletedList string = string(wikiwikimagic.TriangularBulletedList)
			const magicStringHyphenBulletedList     string = string(wikiwikimagic.HyphenBulletedList)

			switch string(element) {
			case magicStringHeading:
				element = ""

				var i int
				miniloop1: for i < 256 {
					i++

					var size int

					r, size, err := utf8.ReadRune(bufferedreader)
					if 0 < size {
						switch r {
						case '§':
							element = element+internalElement(string(r))
							continue
						case ' ':
							element = element+internalElement(string(r))
							break miniloop1
						default:
							element = element+internalElement(string(r))
							break miniloop1
						}
					}
					if errors.Is(err, io.EOF) {
						io.WriteString(writer, "<p>")
						io.WriteString(writer, string(element))
						io.WriteString(writer, "</p>")
						return nil
					}
					if nil != err {
						return err
					}
				}
			case magicStringRoundBulletedList, magicStringTriangularBulletedList, magicStringHyphenBulletedList:
				element = ""

				var i int
				miniloop2: for i < 256 {
					i++

					r, size, err := utf8.ReadRune(bufferedreader)
					if 0 < size {
						switch r {
						case wikiwikimagic.RoundBulletedList, wikiwikimagic.TriangularBulletedList, wikiwikimagic.HyphenBulletedList:
							element = element+internalElement(string(r))
							continue
						case ' ':
							element = element+internalElement(string(r))
							break miniloop2
						default:
							element = element+internalElement(string(r))
							break miniloop2
						}
					}
					if errors.Is(err, io.EOF) {
						io.WriteString(writer, "<p>")
						io.WriteString(writer, string(element))
						io.WriteString(writer, "</p>")
						return nil
					}
					if nil != err {
						return err
					}
				}
			}
		}

		var block io.ReadCloser = iopsep.NewParagraphReadCloser(bufferedreader)
		if nil == block {
			return errNilBlockReader
		}
		defer block.Close()

		{
			var code string = element.Begin()

			_, err := io.WriteString(writer, code)
			if nil != err {
				return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
			}

		}

		for {
			var line io.ReadCloser = iolsep.NewLineReadCloser(block)
			if nil == line {
				return errNilLineReader
			}
			defer line.Close()

			err = nil
			var size int

			var numRunes int
			for nil == err {
				var r rune

				r, size, err = utf8.ReadRune(line)
				if 0 < size {
					numRunes++

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
			}
			if errors.Is(err,io.EOF) {
				if numRunes <= 0 {
		/////////////////////// BREAK
					break
				}
		/////////////// CONTINUE
				continue
			}
			if nil == err && size <= 0 {
				return errInternalError
			}
			return err
		}

		{
			var code string = element.End()

			_, err := io.WriteString(writer, code)
			if nil != err {
				return erorr.Errorf("wikiwiki: problem writing %q: %w", code, err)
			}
		}

	}

	return nil
}
