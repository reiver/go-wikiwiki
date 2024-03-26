package wikiwiki

import (
	"io"
)

type BlockRenderer interface {
	RenderLine(io.Writer, io.RuneScanner) error
	BeginRender(io.Writer) error
	EndRender(io.Writer) error
}
