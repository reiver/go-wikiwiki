package wikiwikimagic

func isSpacing(r rune) bool {
	switch r {
	case HT, SP, NBSP:
		return true
	default:
		return false
	}
}
