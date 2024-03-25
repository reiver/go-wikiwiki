package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var highLightRenderer wikiwiki.ElementRenderer = internalHighLightRenderer{}

func (internalRenderer) HighLightRenderer() wikiwiki.ElementRenderer {
	return internalHighLightRenderer{}
}

type internalHighLightRenderer struct{}

func (internalHighLightRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalHighLightRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<mark>")
}

func (receiver internalHighLightRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</mark>")
}
