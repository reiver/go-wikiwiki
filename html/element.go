package wikiwikihtml

import (
	"github.com/reiver/go-wikiwiki/magic"
)

const pre string = string(wikiwikimagic.Preformatted)

const listrb string = string(wikiwikimagic.RoundBulletedList)+" "
const listtb string = string(wikiwikimagic.TriangularBulletedList)+" "
const listhb string = string(wikiwikimagic.HyphenBulletedList)

const h1  string = "§ "
const h2  string = "§§ "
const h3  string = "§§§ "
const h4  string = "§§§§ "
const h5  string = "§§§§§ "
const h6  string = "§§§§§§ "
const h7  string = "§§§§§§§ "
const h8  string = "§§§§§§§§ "
const h9  string = "§§§§§§§§§ "
const h10 string = "§§§§§§§§§§ "
const h11 string = "§§§§§§§§§§§ "
const h12 string = "§§§§§§§§§§§§ "

const quotation string = string(wikiwikimagic.Quotation)+" "

type internalElement string

func (receiver internalElement) Begin() string {
	switch string(receiver) {
//@TODO
//	case pre:
//		return `<pre style="line-height:0.125em;">`+"\n"
	case pre:
		return `<pre>`+"\n"
	case listrb:
		return `<ul style="list-style-type:disc">`
	case listtb:
		return `<ul style="list-style-type:&quot‣;&quot;">`
	case listhb:
		return `<ul style="list-style-type:&quot;⁃&quot;">`
	case h1:
		return "<h1>"
	case h2:
		return "<h2>"
	case h3:
		return "<h3>"
	case h4:
		return "<h4>"
	case h5:
		return "<h5>"
	case h6:
		return "<h6>"
	case h7:
		return `<div class="h7">`
	case h8:
		return `<div class="h8">`
	case h9:
		return `<div class="h9">`
	case h10:
		return `<div class="h10">`
	case h11:
		return `<div class="h11">`
	case h12:
		return `<div class="h12">`
	case quotation:
		return "<blockquote>\n"
	default:
		return "<p>\n"
	}
}

func (receiver internalElement) End() string {
	switch string(receiver) {
	case pre:
		return `</pre>`+"\n"
	case listrb, listtb, listhb:
		return "</ul>\n"
	case h1: // U+00A7 Section Sign
		return "</h1>\n"
	case h2:
		return "</h2>\n"
	case h3:
		return "</h3>\n"
	case h4:
		return "</h4>\n"
	case h5:
		return "</h5>\n"
	case h6:
		return "</h6>\n"
	case h7:
		return "</div>"
	case h8:
		return "</div>"
	case h9:
		return "</div>"
	case h10:
		return "</div>"
	case h11:
		return "</div>"
	case h12:
		return "</div>"
	case quotation:
		return "</blockquote>\n"
	default:
		return "</p>\n"
	}
}
