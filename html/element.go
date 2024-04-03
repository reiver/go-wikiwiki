package wikiwikihtml

type internalElement string

func (receiver internalElement) Begin() string {
	switch string(receiver) {
	case ".": // U+00A7 Section Sign
		return `<pre style="line-height:0.125em;">`+"\n"
	case "•", // U+2022 Bullet
	     "‣", // U+2023 Triangular Bullet
	     "⁃": // U+2043 Hyphen Bullet
		return "<ul>\n<li>"
	case "§": // U+00A7 Section Sign
		return "<h1>"
	case "§§":
		return "<h2>"
	case "§§§":
		return "<h3>"
	case "§§§§":
		return "<h4>"
	case "§§§§§":
		return "<h5>"
	case "§§§§§§":
		return "<h6>"
	case "―": // U+2015 Horizontal Bar; i.e., quotation dash.
		return "<blockquote>\n"
	default:
		return "<p>\n"
	}
}

func (receiver internalElement) End() string {
	switch string(receiver) {
	case ".": // U+00A7 Section Sign
		return `</pre>`+"\n"
	case "•", // U+2022 Bullet
	     "‣", // U+2023 Triangular Bullet
	     "⁃": // U+2043 Hyphen Bullet
		return "</ul>\n"
	case "§": // U+00A7 Section Sign
		return "</h1>\n"
	case "§§":
		return "</h2>\n"
	case "§§§":
		return "</h3>\n"
	case "§§§§":
		return "</h4>\n"
	case "§§§§§":
		return "</h5>\n"
	case "§§§§§§":
		return "</h6>\n"
	case "―": // U+2015 Horizontal Bar; i.e., quotation dash.
		return "</blockquote>\n"
	default:
		return "</p>\n"
	}
}

func (receiver internalElement) Buffer() string {
	switch string(receiver) {
	case ".": // U+00A7 Section Sign
		// nothing here
	case "•", // U+2022 Bullet
	     "‣", // U+2023 Triangular Bullet
	     "⁃": // U+2043 Hyphen Bullet
		// nothing here
	case "§": // U+00A7 Section Sign
		// nothing here
	case "§§":
		// nothing here
	case "§§§":
		// nothing here
	case "§§§§":
		// nothing here
	case "§§§§§":
		// nothing here
	case "§§§§§§":
		// nothing here
	case "―": // U+2015 Horizontal Bar; i.e., quotation dash.
		// nothing here
	default:
		return string(receiver)
	}

	return ""
}
