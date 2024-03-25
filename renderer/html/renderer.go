package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

type internalRenderer struct {}

// NewRenderer returns a new wiki text to HTML renderer.
func NewRenderer() wikiwiki.Renderer {
	return internalRenderer{}
}

func (internalRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}
