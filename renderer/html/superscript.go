package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var superScriptRenderer wikiwiki.ElementRenderer = internalSuperScriptRenderer{}

func (internalRenderer) SuperScriptRenderer() wikiwiki.ElementRenderer {
	return internalSuperScriptRenderer{}
}

type internalSuperScriptRenderer struct{}

func (internalSuperScriptRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalSuperScriptRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<sup>")
}

func (receiver internalSuperScriptRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</sup>")
}
