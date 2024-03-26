package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var infoRenderer wikiwiki.BlockRenderer = internalInfoRenderer{}

func (internalRenderer) InfoRenderer() wikiwiki.BlockRenderer {
	return infoRenderer
}

type internalInfoRenderer struct{}

func (internalInfoRenderer) RenderLine(writer io.Writer, runescanner io.RuneScanner) error {
	return wikiwiki.TranscodeLine(writer, renderer, runescanner)
}

func (receiver internalInfoRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<div class="wiki-info" style="color:#004085;background-color:#cce5ff;border-color:#b8daff">`)
}

func (receiver internalInfoRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</div>")
}
