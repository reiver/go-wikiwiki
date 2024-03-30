package wikiwikihtmlrune

import (
	"io"
)

func RenderRuneInAttribute(writer io.Writer, r rune) error {
	switch r {
	case '"':
		return write(writer, "&quot;")
	case '&':
		return write(writer, "&amp;")
	case '\'':
		return write(writer, "&apos;")
	case '<':
		return write(writer, "&lt;")
	case '>':
		return write(writer, "&gt;")
	default:
		return write(writer, string(r))
	}
}
