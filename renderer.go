package wikiwiki

// Renderer is something that renders wiki to something else (such as HTML, ANSI, markdown, etc).
//
// For example, one renderer might transcode wiki text to HTML.
// Another renderer might transcode wiki text to ANSI.
// Etc.
type Renderer interface {
	RenderRune(r rune) error

	// **bold**
	BeginBold() error
	EndBold() error

	// ||highlight||
	BeginHighLight() error
	EndHighLight() error

	// //italics//
	BeginItalics() error
	EndItalics() error

	// ::code:: — for example: ::u+2e3a::
	BeginCode() error
	EndCode() error

	// [link]
	BeginLink() error
	EndLink() error

	// ``monospaced``
	BeginMonoSpaced() error
	EndMonoSpaced() error

	// ‾‾overline‾‾
	BeginOverLine() error
	EndOverLine() error

	// ,,subscript,,
	BeginSubScript() error
	EndSubScript() error

	// ''superscript''
	BeginSuperScript() error
	EndSuperScript() error

	// __underline__
	BeginUnderLine() error
	EndUnderLine() error

	// {name}
	BeginNameValue() error
	EndNameValue() error
}
