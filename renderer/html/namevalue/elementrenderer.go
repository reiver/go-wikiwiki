package wikiwikihtmlnamevalue

import (
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-wikiwiki/renderer"
	"github.com/reiver/go-wikiwiki/renderer/html/rune"
)

const elementbegin string = `<span class="wiki-name-value">`
const elementend   string = "</span>"
const elementname  string = "namevalue"

var ElementRenderer wikiwikirenderer.ElementRenderer = internalElementRenderer{}

type internalElementRenderer struct {}

func (internalElementRenderer) RenderRune(writer io.Writer, r rune) error {
	if nil == writer {
		return errNilWriter
	}

	return wikiwikihtmlrune.RenderRune(writer, r)
}

func (internalElementRenderer) BeginRenderingElement(writer io.Writer) error {
	if nil == writer {
		return errNilWriter
	}

	{
		const str string = elementbegin

		_, err := io.WriteString(writer, str)
		if nil != err {
			return erorr.Errorf("wikiwiki: problem writing %s element beginning (%q): %w", elementname, str, err)
		}
	}

	return nil
}

func (internalElementRenderer) EndRenderingElement(writer io.Writer) error {
	if nil == writer {
		return errNilWriter
	}


	{
		const str string = elementend

		_, err := io.WriteString(writer, str)
		if nil != err {
			return erorr.Errorf("wikiwiki: problem writing %s element ending (%q): %w", elementname, str, err)
		}
	}

	return nil
}
