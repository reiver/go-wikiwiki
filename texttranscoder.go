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

	var begin func(io.Writer)error
	var end func(io.Writer)error
	noop := func(io.Writer) error {
		return nil
	}

	switch r {
	default:
		return renderer.RenderRune(writer, r)
	case '\'': // ''superscript''
		begin = renderer.BeginSuperScript
		end   = renderer.EndSuperScript
	case '(',')':  // ((no-op))
		begin = noop
		end = noop
	case '*':  // **bold**
		begin = renderer.BeginBold
		end   = renderer.EndBold
	case ',':  // ,,subscript,,
		begin = renderer.BeginSubScript
		end   = renderer.EndSubScript
	case '/':  // //italics//
		begin = renderer.BeginItalics
		end   = renderer.EndItalics
	case ':':  // ::smartcode::
		begin = renderer.BeginCode
		end   = renderer.EndCode
	case '[',']':  // [[internal link]]
		begin = renderer.BeginLink
		end   = renderer.EndLink
	case '_':  // __underline__
		begin = renderer.BeginUnderLine
		end   = renderer.EndUnderLine
	case '`':  // ``mononspaced``
		begin = renderer.BeginMonoSpaced
		end   = renderer.EndMonoSpaced
	case '{','}':  // {{name}}
		begin = renderer.BeginNameValue
		end   = renderer.EndNameValue
	case '|':  // ||highlight||
		begin = renderer.BeginHighLight
		end   = renderer.EndHighLight
	case '‾':  // ‾‾overline‾‾
		begin = renderer.BeginOverLine
		end   = renderer.EndOverLine
	}

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
