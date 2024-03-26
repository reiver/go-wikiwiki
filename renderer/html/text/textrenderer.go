package wikiwikihtmltext

import (
	"io"

	"github.com/reiver/go-wikiwiki/renderer"
	"github.com/reiver/go-wikiwiki/renderer/html/bold"
	"github.com/reiver/go-wikiwiki/renderer/html/code"
	"github.com/reiver/go-wikiwiki/renderer/html/highlight"
	"github.com/reiver/go-wikiwiki/renderer/html/italics"
	"github.com/reiver/go-wikiwiki/renderer/html/link"
	"github.com/reiver/go-wikiwiki/renderer/html/monospaced"
	"github.com/reiver/go-wikiwiki/renderer/html/namevalue"
	"github.com/reiver/go-wikiwiki/renderer/html/overline"
	"github.com/reiver/go-wikiwiki/renderer/html/rune"
	"github.com/reiver/go-wikiwiki/renderer/html/subscript"
	"github.com/reiver/go-wikiwiki/renderer/html/superscript"
	"github.com/reiver/go-wikiwiki/renderer/html/underline"
)

var TextRenderer wikiwikirenderer.TextRenderer = internalTextRenderer{}

type internalTextRenderer struct {}

func (internalTextRenderer) RenderRune(writer io.Writer, r rune) error {
	return wikiwikihtmlrune.RenderRune(writer, r)
}

// **bold**
func (internalTextRenderer) BoldRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlbold.ElementRenderer
}

// ||highlight||
func (internalTextRenderer) HighLightRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlhighlight.ElementRenderer
}

// //italics//
func (internalTextRenderer) ItalicsRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlitalics.ElementRenderer
}

// ::code:: — for example: ::u+2e3a:: ::icon:banana.png:: ::image:photo.jpeg::
func (internalTextRenderer) CodeRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlcode.ElementRenderer
}

// [link]
func (internalTextRenderer) LinkRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmllink.ElementRenderer
}

// ``monospaced``
func (internalTextRenderer) MonoSpacedRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlmonospaced.ElementRenderer
}

// {name}
func (internalTextRenderer) NameValueRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlnamevalue.ElementRenderer
}

// ‾‾overline‾‾
func (internalTextRenderer) OverLineRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmloverline.ElementRenderer
}

// ,,subscript,,
func (internalTextRenderer) SubScriptRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlsubscript.ElementRenderer
}

// ''superscript''
func (internalTextRenderer) SuperScriptRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlsuperscript.ElementRenderer
}

// __underline__
func (internalTextRenderer) UnderLineRenderer() wikiwikirenderer.ElementRenderer {
	return wikiwikihtmlunderline.ElementRenderer
}
