package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var alertRenderer wikiwiki.BlockRenderer = internalAlertRenderer{}

func (internalRenderer) AlertRenderer() wikiwiki.BlockRenderer {
	return alertRenderer
}

type internalAlertRenderer struct{}

func (internalAlertRenderer) RenderLine(writer io.Writer, runescanner io.RuneScanner) error {
	return wikiwiki.TranscodeLine(writer, renderer, runescanner)
}

func (receiver internalAlertRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<div class="wiki-alert" style="color:#721c24;background-color:#f8d7da;border-color:#f5c6cb">`)
}

func (receiver internalAlertRenderer) EndRender(writer io.Writer) error {
	 return write(writer, "</div>")
}
