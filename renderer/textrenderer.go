package wikiwikirenderer

type TextRenderer interface {
	RuneRenderer

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
