package paragraph

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInternalError = erorr.Error("wikiwiki internal-error")
	errNilReader     = erorr.Error("wikiwiki: nil reader")
	errNilWriter     = erorr.Error("wikiwiki: nil writer")
)
