package wikiwikihtmlrune

import (
	"io"

	"github.com/reiver/go-wikiwiki/renderer"
)

var RuneRenderer wikiwikirenderer.RuneRenderer = internalRuneRenderer{}

func RenderRune(writer io.Writer, r rune) error {
	switch r {
	case '&':
		return write(writer, "&amp;")
	case '<':
		return write(writer, "&lt;")
	case '>':
		return write(writer, "&gt;")
	default:
		return write(writer, string(r))
	}
}

type internalRuneRenderer struct {}

func (internalRuneRenderer) RenderRune(writer io.Writer, r rune) error {
	return RenderRune(writer, r)
}
