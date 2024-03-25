package wikiwiki

import (
	"io"

	"sourcecode.social/reiver/go-opt"
)

type TextTranscoder interface {
	io.Closer

	InterpretRune(rune)error
}

type internalTextTranscoder struct {
	writer io.Writer
	renderer Renderer

	buffer opt.Optional[rune]

	stack stack[rune]
}

func NewTextTranscoder(writer io.Writer, renderer Renderer) TextTranscoder {
	return &internalTextTranscoder{
		writer:writer,
		renderer:renderer,
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

	renderer := receiver.renderer
	if nil == renderer {
		return errNilRenderer
	}

	return renderer.RenderRune(writer, value)
}

func (receiver *internalTextTranscoder) InterpretRune(r rune) (err error) {
	if nil == receiver {
		return errNilReceiver
	}

	writer := receiver.writer
	if nil == writer {
		return errNilWriter
	}

	renderer := receiver.renderer
	if nil == receiver {
		return errNilRenderer
	}

	{
		rr, something := receiver.buffer.Get()
		if something && rr != r {
			err = renderer.RenderRune(writer, rr)
			if nil != err {
				return err
			}
			receiver.buffer = opt.Nothing[rune]()
		}
	}

	var elementrenderer ElementRenderer = noop

	switch r {
	default:
		return renderer.RenderRune(writer, r)
	case '\'': // ''superscript''
		elementrenderer = renderer.SuperScriptRenderer()
	case '(',')':  // ((no-op))
		elementrenderer = noop
	case '*':  // **bold**
		elementrenderer = renderer.BoldRenderer()
	case ',':  // ,,subscript,,
		elementrenderer = renderer.SubScriptRenderer()
	case '/':  // //italics//
		elementrenderer = renderer.ItalicsRenderer()
	case ':':  // ::smartcode::
		elementrenderer = renderer.CodeRenderer()
	case '[',']':  // [[internal link]]
		elementrenderer = renderer.LinkRenderer()
	case '_':  // __underline__
		elementrenderer = renderer.UnderLineRenderer()
	case '`':  // ``mononspaced``
		elementrenderer = renderer.MonoSpacedRenderer()
	case '{','}':  // {{name}}
		elementrenderer = renderer.NameValueRenderer()
	case '|':  // ||highlight||
		elementrenderer = renderer.HighLightRenderer()
	case '‾':  // ‾‾overline‾‾
		elementrenderer = renderer.OverLineRenderer()
	}

	var begin func(io.Writer)error = elementrenderer.BeginRender
	var end func(io.Writer)error = elementrenderer.EndRender

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
