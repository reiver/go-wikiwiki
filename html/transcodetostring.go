package wikiwikihtml

import (
	"strings"

	"github.com/reiver/go-htmlescape"
)

func TransCodeToString(wikitext string) string {
	result, err := TransCodeToStringError(wikitext)
	if nil != err {
		return htmlescape.String(wikitext)
	}

	return result
}

func TransCodeToStringError(wikitext string) (string, error) {
	var dst strings.Builder

	err := Transcode(&dst, strings.NewReader(wikitext))
	if nil != err {
		var nada string
		return nada, err
	}

	return dst.String(), nil
}
