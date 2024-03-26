package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var subScriptRenderer wikiwiki.ElementRenderer = internalSubScriptRenderer{}

func (internalRenderer) SubScriptRenderer() wikiwiki.ElementRenderer {
	return subScriptRenderer
}

type internalSubScriptRenderer struct{}

func (internalSubScriptRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalSubScriptRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<sub>")
}

func (receiver internalSubScriptRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</sub>")
}
