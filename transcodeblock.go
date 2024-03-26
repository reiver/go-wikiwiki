package wikiwiki

import (
	"errors"
	"io"

	"sourcecode.social/reiver/go-eol"
	"sourcecode.social/reiver/go-erorr"
)

func TranscodeBlock(writer io.Writer, renderer Renderer, runescanner io.RuneScanner) (err error) {
	if nil == writer {
		return errNilWriter
	}
	if nil == renderer {
		return errNilRenderer
	}
	if nil == runescanner {
		return errNilReader
	}

	for {
		err = TranscodeNonEmptyLine(writer, renderer, runescanner)

		if errors.Is(err, ErrEmptyLine) {
			endofline, _, err := eol.ReadEOL(runescanner)
			if nil != err {
				return erorr.Errorf("wikwiki: problem reading end-of-line character(s): %w", err)
			}

			for _, r := range endofline {
				e := renderer.RenderRune(writer, r)
				if nil != e {
					return e
				}
			}

	/////////////// BREAK
			break
		}

		if errors.Is(err, io.EOF) {
			return nil
		}

		if nil != err {
			return err
		}
	}

	return nil
}
