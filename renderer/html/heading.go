package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var headingRenderer wikiwiki.BlockRenderer = internalHeadingRenderer{}

func (internalRenderer) HeadingRenderer() wikiwiki.BlockRenderer {
	return headingRenderer
}

type internalHeadingRenderer struct{}

func (internalHeadingRenderer) RenderLine(writer io.Writer, runescanner io.RuneScanner) error {
	return wikiwiki.TranscodeLine(writer, renderer, runescanner)
}

func (receiver internalHeadingRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, "<h1>")
}

func (receiver internalHeadingRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</h1>")
}
