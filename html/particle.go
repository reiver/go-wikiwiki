package wikiwikihtml

import (
	"github.com/reiver/go-unicode"
)

type internalParticle rune

func (receiver internalParticle) Render() (string, bool) {
	var r rune = rune(receiver)

	switch r {
	case unicode.LS:
		return "<br />\n", true
//@TODO
/*
	case '■': // U+25A0 Black Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:1em;`+
				`height:1em;`+
				`background-color:#4f1802;`+
			`"></div>`, true
	case '□': // U+25A1 White Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:1em;`+
				`height:1em;`+
				`background-color:#FF7F00;`+
			`"></div>`, true
	case '▪': // U+25AA Black Small Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:0.5em;`+
				`height:0.5em;`+
				`background-color:#4f1802;`+
			`"></div>`, true
	case '▫': // U+25AB Black Small Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:0.5em;`+
				`height:0.5em;`+
				`background-color:#FF7F00;`+
			`"></div>`, true
	case '⬝': // U+2B1D Black Very Small Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:0.25em;`+
				`height:0.25em;`+
				`background-color:#4f1802;`+
			`"></div>`, true
	case '⬞': // U+2B1D White Very Small Square
		return `<div style="`+
				`display:inline-block;`+
				`margin:0;`+
				`padding:0;`+
				`width:0.25em;`+
				`height:0.25em;`+
				`background-color:#FF7F00;`+
			`"></div>`, true

*/
	case '•', // U+2022 Bullet
	     '‣', // U+2023 Triangular Bullet
	     '⁃': // U+2043 Hyphen Bullet
		return "<li>", true
	default:
		return "", false
	}
}
