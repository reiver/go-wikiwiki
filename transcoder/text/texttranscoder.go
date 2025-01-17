package wikiwikitexttranscoder

import (
	"io"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-wikiwiki/internal/stack"
	"github.com/reiver/go-wikiwiki/renderer"
	"github.com/reiver/go-wikiwiki/transcoder"
)

type internalTextTranscoder struct {
	writer io.Writer
	textrenderer wikiwikirenderer.TextRenderer

	buffer opt.Optional[rune]

	stack stack.Stack[rune]
}

func NewTextTranscoder(writer io.Writer, textrenderer wikiwikirenderer.TextRenderer) wikiwikitranscoder.TextTranscoder {
	return &internalTextTranscoder{
		writer:writer,
		textrenderer:textrenderer,
	}
}

func (receiver *internalTextTranscoder) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	value, found := receiver.buffer.Get()
	if !found {
		return nil
	}

	writer := receiver.writer
	if nil == writer {
		return errNilWriter
	}

	textrenderer := receiver.textrenderer
	if nil == textrenderer {
		return errNilTextRenderer
	}

	return textrenderer.RenderRune(writer, value)
}

func (receiver *internalTextTranscoder) InterpretRune(r rune) (err error) {
	if nil == receiver {
		return errNilReceiver
	}

	writer := receiver.writer
	if nil == writer {
		return errNilWriter
	}

	textrenderer := receiver.textrenderer
	if nil == textrenderer {
		return errNilTextRenderer
	}

	{
		rr, something := receiver.buffer.Get()
		if something && rr != r {
			err = textrenderer.RenderRune(writer, rr)
			if nil != err {
				return err
			}
			receiver.buffer = opt.Nothing[rune]()
		}
	}

	var elementrenderer wikiwikirenderer.ElementRenderer = wikiwikirenderer.NoOpElementRenderer

	switch r {
	default:
		return textrenderer.RenderRune(writer, r)
	case '\'': // ''superscript''
		elementrenderer = textrenderer.SuperScriptRenderer()
	case '(',')':  // ((no-op))
		elementrenderer = wikiwikirenderer.NoOpElementRenderer
	case '*':  // **bold**
		elementrenderer = textrenderer.BoldRenderer()
	case ',':  // ,,subscript,,
		elementrenderer = textrenderer.SubScriptRenderer()
	case '/':  // //italics//
		elementrenderer = textrenderer.ItalicsRenderer()
	case ':':  // ::smartcode::
		elementrenderer = textrenderer.CodeRenderer()
	case '[',']':  // [[internal link]]
		elementrenderer = textrenderer.LinkRenderer()
	case '_':  // __underline__
		elementrenderer = textrenderer.UnderLineRenderer()
	case '`':  // ``mononspaced``
		elementrenderer = textrenderer.MonoSpacedRenderer()
	case '{','}':  // {{name}}
		elementrenderer = textrenderer.NameValueRenderer()
	case '|':  // ||highlight||
		elementrenderer = textrenderer.HighLightRenderer()
	case '‾':  // ‾‾overline‾‾
		elementrenderer = textrenderer.OverLineRenderer()
	}

	var begin func(io.Writer)error = elementrenderer.BeginRenderingElement
	var end func(io.Writer)error = elementrenderer.EndRenderingElement

	if opt.Something(r) == receiver.buffer {
		receiver.buffer = opt.Nothing[rune]()
		switch {
		case (')' != r && ']' != r && '}' != r) && receiver.stack.TopEqual(r):
			receiver.stack.Pop()
			return end(writer)
		case ')' == r && receiver.stack.TopEqual('(') ||
		     ']' == r && receiver.stack.TopEqual('[') ||
		     '}' == r && receiver.stack.TopEqual('{'):
			receiver.stack.Pop()
			return end(writer)
		default:
			receiver.stack.Push(r)
			return begin(writer)
		}
	} else {
		receiver.buffer = opt.Something(r)
		return nil
	}
}
