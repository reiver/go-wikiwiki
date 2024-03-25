package wikiwiki

import (
	"io"
)

var noop ElementRenderer = noopElementRenderer{}

type noopElementRenderer struct {}

func (noopElementRenderer) RenderRune(io.Writer, rune) error {
	return nil
}

func (noopElementRenderer) BeginRender(io.Writer) error {
	return nil
}

func (noopElementRenderer) EndRender(io.Writer) error {
	return nil
}

