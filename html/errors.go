package wikiwikihtml

import (
	"github.com/reiver/go-erorr"
)

const (
	errInternalError     = erorr.Error("wikiwiki: internal error")
	errNilBlockReader    = erorr.Error("wikiwiki: nil block-reader")
	errNilBufferedReader = erorr.Error("wikiwiki: nil buffered-reader")
	errNilFile           = erorr.Error("wikiwiki: nil file")
	errNilLineReader     = erorr.Error("wikiwiki: nil line-reader")
	errNilPipeReader     = erorr.Error("wikiwiki: nil pipe-reader")
	errNilPipeWriter     = erorr.Error("wikiwiki: nil pipe-writer")
	errNilReader         = erorr.Error("wikiwiki: nil reader")
	errNilReceiver       = erorr.Error("wikiwiki: nil receiver")
	errNilRuneScanner    = erorr.Error("wikiwiki: nil rune-scanner")
	errNilWriter         = erorr.Error("wikiwiki: nil writer")
)
