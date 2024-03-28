package wikiwikitexttranscoder

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilReceiver     = erorr.Error("wikiwiki: nil receiver")
	errNilTextRenderer = erorr.Error("wikiwiki: nil text-renderer")
	errNilWriter       = erorr.Error("wikiwiki: nil writer")
)
