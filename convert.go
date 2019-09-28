package rungo

func convertDigitByte(ch byte) (digit int, ok bool) {
	if (ch >= '0') && (ch <= '9') {
		digit = int(ch - '0')
		ok = true
		return
	}
	return -1, false
}
