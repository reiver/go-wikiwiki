package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

type internalRenderer struct {
	writer io.Writer
}

// NewRenderer returns a new wiki text to HTML renderer.
func NewRenderer(writer io.Writer) wikiwiki.Renderer {
	return internalRenderer{
		writer:writer,
	}
}

func (receiver internalRenderer) write(s string) error {
	writer := receiver.writer
	if nil == writer {
		return errNilWriter
	}

	_, err := io.WriteString(writer, s)
	return err
}

func (receiver internalRenderer) BeginBold() error {
	return receiver.write("<b>")

}

func (receiver internalRenderer) EndBold() error {
	return receiver.write("</b>")
}

func (receiver internalRenderer) BeginCode() error {
	return receiver.write(`<span class="wiki-code">`)

}

func (receiver internalRenderer) EndCode() error {
	return receiver.write("</span>")
}

func (receiver internalRenderer) BeginHighLight() error {
	return receiver.write("<mark>")
}

func (receiver internalRenderer) EndHighLight() error {
	return receiver.write("</mark>")
}

func (receiver internalRenderer) BeginItalics() error {
	return receiver.write("<i>")

}

func (receiver internalRenderer) EndItalics() error {
	return receiver.write("</i>")

}

func (receiver internalRenderer) BeginLink() error {
	return receiver.write(`<span class="wiki-link">`)

}

func (receiver internalRenderer) EndLink() error {
	return receiver.write(`</span>`+`<script>`+`(function(){`+
		`e=document.querySelector(".wiki-link");`+
		`n=document.createElement("a");`+
		`t=e.innerText;`+
		`n.href=t;`+
		`n.innerText=t;`+
		`e.replaceWith(n)`+
		`})()`+`</script>`)

}

func (receiver internalRenderer) BeginMonoSpaced() error {
	return receiver.write("<tt>")

}

func (receiver internalRenderer) EndMonoSpaced() error {
	return receiver.write("</tt>")

}

func (receiver internalRenderer) BeginOverLine() error {
	return receiver.write(`<span style="text-decoration:overline">`)

}

func (receiver internalRenderer) EndOverLine() error {
	return receiver.write("</span>")

}

func (receiver internalRenderer) BeginSubScript() error {
	return receiver.write("<sub>")

}

func (receiver internalRenderer) EndSubScript() error {
	return receiver.write("</sub>")

}

func (receiver internalRenderer) BeginSuperScript() error {
	return receiver.write("<sup>")

}

func (receiver internalRenderer) EndSuperScript() error {
	return receiver.write("</sup>")

}

func (receiver internalRenderer) BeginUnderLine() error {
	return receiver.write("<u>")

}

func (receiver internalRenderer) EndUnderLine() error {
	return receiver.write("</u>")

}

func (receiver internalRenderer) BeginNameValue() error {
	return receiver.write(`<span class="wiki-name-value">`)

}

func (receiver internalRenderer) EndNameValue() error {
	return receiver.write(`</span>`+`<script>`+`(function(){`+
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

func (receiver internalRenderer) RenderRune(r rune) error {
	switch r {
	case '&':
		return receiver.write("&amp;")
	case '<':
		return receiver.write("&lt;")
	case '>':
		return receiver.write("&gt;")
	default:
		return receiver.write(string(r))
	}
}
