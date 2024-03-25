package wikiwiki

import (
	"io"
)

// ElementRenderer is a renderer a single type of element.
//
// For example, it might just handle rendering a bold element.
// Or, also for example, it might just handle rendering an italics element.
// Etc.
//
// Typically, one gets an ElementRenderer from Renderer.
type ElementRenderer interface {
	RenderRune(io.Writer, rune) error
	BeginRender(io.Writer) error
	EndRender(io.Writer) error
}
