package wikiwikihtml

import (
	"io"
)

func renderrune(writer io.Writer, r rune) error {
	switch r {
	case '&':
		return write(writer, "&amp;")
	case '<':
		return write(writer, "&lt;")
	case '>':
		return write(writer, "&gt;")
	default:
		return write(writer, string(r))
	}
}
