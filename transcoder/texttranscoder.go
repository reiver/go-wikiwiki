package wikiwikitranscoder

import (
	"io"
)

type TextTranscoder interface {
	io.Closer

	InterpretRune(rune)error
}
