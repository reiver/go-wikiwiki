package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

var nameValueRenderer wikiwiki.ElementRenderer = internalNameValueRenderer{}

func (internalRenderer) NameValueRenderer() wikiwiki.ElementRenderer {
	return nameValueRenderer
}

type internalNameValueRenderer struct{}

func (internalNameValueRenderer) RenderRune(writer io.Writer, r rune) error {
	return renderrune(writer, r)
}

func (receiver internalNameValueRenderer)  BeginRender(writer io.Writer) error {
	 return write(writer, `<span class="wiki-name-value">`)
}

func (receiver internalNameValueRenderer) EndRender(writer io.Writer) error {
	 return write(writer, `</span>`+`<script>`+`(function(){`+
		`e=document.querySelector(".wiki-name-value");`+
		`m=document.getElementById("local."+e.innerText);`+
		`if(m){`+
		`t=m.getAttribute("content");`+
		`n=document.createElement("span");`+
		`n.innerText=t;`+
		`e.replaceWith(n)`+
		`}`+
		`})()`+`</script>`)
}
