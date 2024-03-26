package wikiwiki

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilReader         = erorr.Error("wikiwiki: nil reader")
	errNilReceiver       = erorr.Error("wikiwiki: nil receiver")
	errNilRenderer       = erorr.Error("wikiwiki: nil renderer")
	errNilTextTranscoder = erorr.Error("wikiwiki: nil text-transcoder")
	errNilWriter         = erorr.Error("wikiwiki: nil writer")
)

const (
	ErrEmptyLine = erorr.Error("empty line")
)
