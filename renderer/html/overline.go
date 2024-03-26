package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var overLineRenderer wikiwiki.ElementRenderer = internalOverLineRenderer{}

func (internalRenderer) OverLineRenderer() wikiwiki.ElementRenderer {
	return overLineRenderer
}

type internalOverLineRenderer struct{}

func (internalOverLineRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalOverLineRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<span style="text-decoration:overline">`)
}

func (receiver internalOverLineRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</span>")
}
