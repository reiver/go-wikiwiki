package wikiwikihtml

import (
	"io"
)

func write(writer io.Writer, s string) error {
	if nil == writer {
		return errNilWriter
	}

	_, err := io.WriteString(writer, s)
	return err
}
