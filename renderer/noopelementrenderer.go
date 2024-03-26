package wikiwikirenderer

import (
	"io"
)

var NoOpElementRenderer ElementRenderer = noopElementRenderer{}

type noopElementRenderer struct {}

func (noopElementRenderer) RenderRune(io.Writer, rune) error {
	return nil
}

func (noopElementRenderer) BeginRenderingElement(io.Writer) error {
	return nil
}

func (noopElementRenderer) EndRenderingElement(io.Writer) error {
	return nil
}

