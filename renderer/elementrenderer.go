package wikiwikirenderer

import (
	"io"
)

type ElementRenderer interface {
	RuneRenderer
	BeginRenderingElement(io.Writer) error
	EndRenderingElement(io.Writer) error
}
