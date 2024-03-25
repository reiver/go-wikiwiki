package wikiwiki

import (
	"io"
)

type ElementRenderer interface {
	RenderRune(io.Writer, rune) error
	BeginRender(io.Writer) error
	EndRender(io.Writer) error
}
