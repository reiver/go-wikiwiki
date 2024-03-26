package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var italicsRenderer wikiwiki.ElementRenderer = internalItalicsRenderer{}

func (internalRenderer) ItalicsRenderer() wikiwiki.ElementRenderer {
	return italicsRenderer
}

type internalItalicsRenderer struct{}

func (internalItalicsRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalItalicsRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<i>")
}

func (receiver internalItalicsRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</i>")
}
