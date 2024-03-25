package wikiwikihtml

import (
	"io"

	"github.com/reiver/go-wikiwiki"
)

type internalRenderer struct {}

// NewRenderer returns a new wiki text to HTML renderer.
func NewRenderer() wikiwiki.Renderer {
	return internalRenderer{}
}

func (receiver internalRenderer) write(writer io.Writer, s string) error {
	if nil == writer {
		return errNilWriter
	}

	_, err := io.WriteString(writer, s)
	return err
}

func (receiver internalRenderer) BeginBold(writer io.Writer) error {
	return receiver.write(writer, "<b>")

}

func (receiver internalRenderer) EndBold(writer io.Writer) error {
	return receiver.write(writer, "</b>")
}

func (receiver internalRenderer) BeginCode(writer io.Writer) error {
	return receiver.write(writer, `<span class="wiki-code">`)

}

func (receiver internalRenderer) EndCode(writer io.Writer) error {
	return receiver.write(writer, "</span>")
}

func (receiver internalRenderer) BeginHighLight(writer io.Writer) error {
	return receiver.write(writer, "<mark>")
}

func (receiver internalRenderer) EndHighLight(writer io.Writer) error {
	return receiver.write(writer, "</mark>")
}

func (receiver internalRenderer) BeginItalics(writer io.Writer) error {
	return receiver.write(writer, "<i>")

}

func (receiver internalRenderer) EndItalics(writer io.Writer) error {
	return receiver.write(writer, "</i>")

}

func (receiver internalRenderer) BeginLink(writer io.Writer) error {
	return receiver.write(writer, `<span class="wiki-link">`)

}

func (receiver internalRenderer) EndLink(writer io.Writer) error {
	return receiver.write(writer, `</span>`+`<script>`+`(function(){`+
		`e=document.querySelector(".wiki-link");`+
		`n=document.createElement("a");`+
		`t=e.innerText;`+
		`n.href=t;`+
		`n.innerText=t;`+
		`e.replaceWith(n)`+
		`})()`+`</script>`)

}

func (receiver internalRenderer) BeginMonoSpaced(writer io.Writer) error {
	return receiver.write(writer, "<tt>")

}

func (receiver internalRenderer) EndMonoSpaced(writer io.Writer) error {
	return receiver.write(writer, "</tt>")

}

func (receiver internalRenderer) BeginOverLine(writer io.Writer) error {
	return receiver.write(writer, `<span style="text-decoration:overline">`)

}

func (receiver internalRenderer) EndOverLine(writer io.Writer) error {
	return receiver.write(writer, "</span>")

}

func (receiver internalRenderer) BeginSubScript(writer io.Writer) error {
	return receiver.write(writer, "<sub>")

}

func (receiver internalRenderer) EndSubScript(writer io.Writer) error {
	return receiver.write(writer, "</sub>")

}

func (receiver internalRenderer) BeginSuperScript(writer io.Writer) error {
	return receiver.write(writer, "<sup>")

}

func (receiver internalRenderer) EndSuperScript(writer io.Writer) error {
	return receiver.write(writer, "</sup>")

}

func (receiver internalRenderer) BeginUnderLine(writer io.Writer) error {
	return receiver.write(writer, "<u>")

}

func (receiver internalRenderer) EndUnderLine(writer io.Writer) error {
	return receiver.write(writer, "</u>")

}

func (receiver internalRenderer) BeginNameValue(writer io.Writer) error {
	return receiver.write(writer, `<span class="wiki-name-value">`)

}

func (receiver internalRenderer) EndNameValue(writer io.Writer) error {
	return receiver.write(writer, `</span>`+`<script>`+`(function(){`+
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

func (receiver internalRenderer) RenderRune(writer io.Writer, r rune) error {
	switch r {
	case '&':
		return receiver.write(writer, "&amp;")
	case '<':
		return receiver.write(writer, "&lt;")
	case '>':
		return receiver.write(writer, "&gt;")
	default:
		return receiver.write(writer, string(r))
	}
}
