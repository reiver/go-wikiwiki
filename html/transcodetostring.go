package wikiwikihtml

import (
	"strings"

	"github.com/reiver/go-htmlescape"
)

func TransCodeToString(wikitext string) string {
	var dst strings.Builder

	err := Transcode(&dst, strings.NewReader(wikitext))
	if nil != err {
		return htmlescape.String(wikitext)
	}

	return dst.String()
}
