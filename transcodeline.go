package wikiwiki

import (
	"errors"
	"io"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-eol"
)

// TranscodeLine transcodes a line of wiki text.
func TranscodeLine(writer io.Writer, renderer Renderer, runescanner io.RuneScanner) (err error) {
	err = TranscodeNonEmptyLine(writer, renderer, runescanner)

	if errors.Is(err, ErrEmptyLine) {
		return nil
	}
	if errors.Is(err, io.EOF) {
		return err
	}
	if nil != err {
		return err
	}

	return nil
}

// TranscodeNonEmptyLine transcodes a non-empty line of wiki text.
// If the line is empty it returns wikiwiki.ErrEmptyLine.
func TranscodeNonEmptyLine(writer io.Writer, renderer Renderer, runescanner io.RuneScanner) (err error) {
	if nil == writer {
		return errNilWriter
	}
	if nil == renderer {
		return errNilRenderer
	}
	if nil == runescanner {
		return errNilReader
	}

	{
		var r rune

		r, _, err = runescanner.ReadRune()

		{
			e := runescanner.UnreadRune()
			if nil != err {
				return e
			}
		}

		if eol.IsEOL(r) {
			return ErrEmptyLine
		}
	}

	var texttranscoder TextTranscoder = NewTextTranscoder(writer, renderer)
	if nil == texttranscoder {
		return errNilTextTranscoder
	}

	defer func(){
		err = texttranscoder.Close()
	}()

	for {
		var r rune
		var n int

		r, n, err = runescanner.ReadRune()
		if 0 < n {
			if eol.IsEOL(r) {
				e := runescanner.UnreadRune()
				if nil != err {
					return e
				}
	/////////////////////// BREAK
				break
			}
			e := texttranscoder.InterpretRune(r)
			if nil != e {
				return e
			}
		}
		if errors.Is(err, io.EOF) {
			return err
		}
		if nil != err {
			return erorr.Errorf("wikwiki: problem reading rune: %w", err)
		}
	}

	{
		endofline, _, err := eol.ReadEOL(runescanner)
		if nil != err {
			return erorr.Errorf("wikwiki: problem reading end-of-line character(s): %w", err)
		}

		for _, r := range endofline {
			e := texttranscoder.InterpretRune(r)
			if nil != e {
				return e
			}
		}
	}

	return nil
}
