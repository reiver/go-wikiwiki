package wikiwikihtml

import (
	"errors"
	"io"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-wikiwiki/renderer"
	"github.com/reiver/go-wikiwiki/renderer/html/text"
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

	var textrenderer wikiwikirenderer.TextRenderer = wikiwikihtmltext.TextRenderer
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
		var r rune

		r, _, err = utf8.ReadRune(reader)
		if errors.Is(err, io.EOF) {
	/////////////// BREAK
			break
		}
		if nil != err {

		}

		err = texttranscoder.InterpretRune(r)
		if nil != err {
			return erorr.Errorf("wikiwiki: text-transcoder had trouble interpretting rune %q (%U): %w", r, r, err)
		}
	}

	return nil
}
