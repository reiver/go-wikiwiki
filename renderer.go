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
	BeginBold(io.Writer) error
	EndBold(io.Writer) error

	// ||highlight||
	BeginHighLight(io.Writer) error
	EndHighLight(io.Writer) error

	// //italics//
	BeginItalics(io.Writer) error
	EndItalics(io.Writer) error

	// ::code:: — for example: ::u+2e3a:: ::icon:banana.png:: ::image:photo.jpeg::
	BeginCode(io.Writer) error
	EndCode(io.Writer) error

	// [link]
	BeginLink(io.Writer) error
	EndLink(io.Writer) error

	// ``monospaced``
	BeginMonoSpaced(io.Writer) error
	EndMonoSpaced(io.Writer) error

	// {name}
	BeginNameValue(io.Writer) error
	EndNameValue(io.Writer) error

	// ‾‾overline‾‾
	BeginOverLine(io.Writer) error
	EndOverLine(io.Writer) error

	// ,,subscript,,
	BeginSubScript(io.Writer) error
	EndSubScript(io.Writer) error

	// ''superscript''
	BeginSuperScript(io.Writer) error
	EndSuperScript(io.Writer) error

	// __underline__
	BeginUnderLine(io.Writer) error
	EndUnderLine(io.Writer) error
}
