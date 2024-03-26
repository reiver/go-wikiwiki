package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var renderer wikiwiki.Renderer = internalRenderer{}

type internalRenderer struct {}

// NewRenderer returns a new wiki text to HTML renderer.
func NewRenderer() wikiwiki.Renderer {
	return renderer
}

func (internalRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}
