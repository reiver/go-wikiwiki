package wikiwiki

import (
	"io"
)

// Renderer is something that renders wiki to something else (such as HTML, ANSI, markdown, etc).
//
// For example, one renderer might transcode wiki text to HTML.
// Another renderer might transcode wiki text to ANSI.
// Etc.
type Renderer interface {
	RenderRune(io.Writer, rune) error

	// **bold**
	BoldRenderer() ElementRenderer

	// ||highlight||
	HighLightRenderer() ElementRenderer

	// //italics//
	ItalicsRenderer() ElementRenderer

	// ::code:: — for example: ::u+2e3a:: ::icon:banana.png:: ::image:photo.jpeg::
	CodeRenderer() ElementRenderer

	// [link]
	LinkRenderer() ElementRenderer

	// ``monospaced``
	MonoSpacedRenderer() ElementRenderer

	// {name}
	NameValueRenderer() ElementRenderer

	// ‾‾overline‾‾
	OverLineRenderer() ElementRenderer

	// ,,subscript,,
	SubScriptRenderer() ElementRenderer

	// ''superscript''
	SuperScriptRenderer() ElementRenderer

	// __underline__
	UnderLineRenderer() ElementRenderer
}
