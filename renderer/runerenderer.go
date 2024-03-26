package wikiwikirenderer

import (
	"io"
)

type RuneRenderer interface {
	RenderRune(io.Writer, rune) error
}
