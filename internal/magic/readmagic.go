package wikiwikimagic

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"

	. "github.com/reiver/go-wikiwiki/internal/eol"
)

func ReadMagic(reader io.Reader) (n int, err error) {

	if nil == reader {
		return n, errNilReader
	}

	var buffer [len(Magic)]byte
	{
		num, err := io.ReadFull(reader, buffer[:])
		n += num
		if nil != err {
			return n, erorr.Errorf("wikwiki: problem reading magic: %w", err)
		}
		{
			const expected int = len(Magic)

			if expected != num {
				return n, erorr.Errorf("wikiwiki: bad read — expected to read %d bytes but only read %d bytes", expected, num)
			}
		}
	}

	{
		if Magic != string(buffer[:]) {
			return n, erorr.Errorf("wikiwiki: bad magic — expected first %d bytes to be %q but actually was %q", len(Magic), Magic, buffer)
		}
	}

	var r rune
	{
		var num int
		var err error

		r, num, err = utf8.ReadRune(reader)
		n += num
		if io.EOF == err {
			return n, nil
		}
		if nil != err {
			return n, erorr.Errorf("wikwiki: problem reading spacing or end-of-line character: %w", err)
		}

	}

	{
		const NBSP rune = '\u00A0'

		switch {
		case IsEOL(r): // inter-line spacing (eol)
			return n, nil
		case isSpacing(r): // intra-line spacing
			// Nothing here.
		default:
			return n, erorr.Errorf("wikiwiki: expected spacing or end-of-line character but instead got %q (%U)", r, r)
		}
	}

	// ignore everything else on the line/
	loop: for {
		var num int
		var err error

		r, num, err = utf8.ReadRune(reader)
		n += num
		if io.EOF == err {
			return n, nil
		}
		if nil != err {
			return n, erorr.Errorf("wikwiki: problem reading character: %w", err)
		}
		if IsEOL(r) { // inter-line spacing (eol)
			break loop
		}
	}

	return n, nil
}
