package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var codeRenderer wikiwiki.ElementRenderer = internalCodeRenderer{}

func (internalRenderer) CodeRenderer() wikiwiki.ElementRenderer {
	return codeRenderer
}

type internalCodeRenderer struct{}

func (internalCodeRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalCodeRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<span class="wiki-code">`)
}

func (receiver internalCodeRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</span>")
}
