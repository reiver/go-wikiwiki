package paragraph

import (
	"github.com/reiver/go-erorr"
)

const (
	errInternalError = erorr.Error("wikiwiki internal-error")
	errNilReader     = erorr.Error("wikiwiki: nil reader")
	errNilWriter     = erorr.Error("wikiwiki: nil writer")
)
