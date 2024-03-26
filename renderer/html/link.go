package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var linkRenderer wikiwiki.ElementRenderer = internalLinkRenderer{}

func (internalRenderer) LinkRenderer() wikiwiki.ElementRenderer {
	return linkRenderer
}

type internalLinkRenderer struct{}

func (internalLinkRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalLinkRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<span class="wiki-link">`)
}

func (receiver internalLinkRenderer) EndRender(writer io.Writer) error {
	 return write(writer, `</span>`+`<script>`+`(function(){`+
		`e=document.querySelector(".wiki-link");`+
		`n=document.createElement("a");`+
		`t=e.innerText;`+
		`n.href=t;`+
		`n.innerText=t;`+
		`e.replaceWith(n)`+
		`})()`+`</script>`)
}
