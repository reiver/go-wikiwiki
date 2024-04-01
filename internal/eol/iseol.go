package wikiwikieol

func IsEOL(r rune) bool {
	switch r {
	case LF, VT, FF, CR, NEL, LS, PS:
		return true
	default:
		return false
	}
}
