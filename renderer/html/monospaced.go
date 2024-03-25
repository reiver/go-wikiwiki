package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var monoSpacedRenderer wikiwiki.ElementRenderer = internalMonoSpacedRenderer{}

func (internalRenderer) MonoSpacedRenderer() wikiwiki.ElementRenderer {
	return internalMonoSpacedRenderer{}
}

type internalMonoSpacedRenderer struct{}

func (internalMonoSpacedRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalMonoSpacedRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<tt>")
}

func (receiver internalMonoSpacedRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</tt>")
}
