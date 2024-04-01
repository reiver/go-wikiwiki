package wikiwikihtml

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInternalError  = erorr.Error("wikiwiki: internal error")
	errNilBlockReader = erorr.Error("wikiwiki: nil block-reader")
	errNilFile        = erorr.Error("wikiwiki: nil file")
	errNilPipeReader  = erorr.Error("wikiwiki: nil pipe-reader")
	errNilPipeWriter  = erorr.Error("wikiwiki: nil pipe-writer")
	errNilReader      = erorr.Error("wikiwiki: nil reader")
	errNilReceiver    = erorr.Error("wikiwiki: nil receiver")
	errNilWriter      = erorr.Error("wikiwiki: nil writer")
)
