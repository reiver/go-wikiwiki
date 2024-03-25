package wikiwiki

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"
)

// TranscodeText transcodes wiki text in 'reader' to something else.
// What that something else is depends on what 'renderer' is used.
// For example, one renderer might transcode wiki text into HTML;
// another renderer might transcode wiki text in ANSI;
// etc.
func TranscodeText(renderer Renderer, reader io.Reader) (err error) {
	if nil == renderer {
		return errNilRenderer
	}
	if nil == reader {
		return errNilReader
	}

	var texttranscoder TextTranscoder = NewTextTranscoder(renderer)
	if nil == texttranscoder {
		return errNilTextTranscoder
	}

	defer func(){
		err = texttranscoder.Close()
	}()

	for {
		var r rune
		var n int

		r, n, err = utf8.ReadRune(reader)
		if 0 < n {
			e := texttranscoder.InterpretRune(r)
			if nil != e {
				return e
			}
		}
		if io.EOF == err {
			return nil
		}
		if nil != err {
			return erorr.Errorf("wikwiki: problem reading rune: %w", err)
		}
	}
}
