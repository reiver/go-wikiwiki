package wikiwiki

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInternalError     = erorr.Error("wikiwiki: internal error")
	errNilReader         = erorr.Error("wikiwiki: nil reader")
	errNilReceiver       = erorr.Error("wikiwiki: nil receiver")
	errNilRenderer       = erorr.Error("wikiwiki: nil renderer")
	errNilTextTranscoder = erorr.Error("wikiwiki: nil text-transcoder")
	errNilWriter         = erorr.Error("wikiwiki: nil writer")
)
