package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var underLineRenderer wikiwiki.ElementRenderer = internalUnderLineRenderer{}

func (internalRenderer) UnderLineRenderer() wikiwiki.ElementRenderer {
	return underLineRenderer
}

type internalUnderLineRenderer struct{}

func (internalUnderLineRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalUnderLineRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<u>")
}

func (receiver internalUnderLineRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</u>")
}
