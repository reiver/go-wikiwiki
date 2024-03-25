package wikiwiki

// Renderer is something that renders wiki to something else (such as HTML, ANSI, markdown, etc).
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
	BeginValue() error
	EndValue() error
}
