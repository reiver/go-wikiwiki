package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var boldRenderer wikiwiki.ElementRenderer = internalBoldRenderer{}

func (internalRenderer) BoldRenderer() wikiwiki.ElementRenderer {
	return boldRenderer
}

type internalBoldRenderer struct{}

func (internalBoldRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalBoldRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<b>")
}

func (receiver internalBoldRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</b>")
}
