package rungo

import (
	"bytes"
	"os/exec"
)

func parseVersionBytes(verBytes []byte) (majorVer, minorVer int, outputText string, err error) {
	state := 0
	for idx, ch := range verBytes {
		switch state {
		case 0:
			if digit, ok := convertDigitByte(ch); ok {
				majorVer = digit
				state = 1
			}
		case 1:
			if digit, ok := convertDigitByte(ch); ok {
				majorVer = majorVer*10 + digit
			} else if ch == '.' {
				state = 2
			} else {
				state = 0
			}
		case 2:
			if digit, ok := convertDigitByte(ch); ok {
				minorVer = minorVer*10 + digit
			} else if bytes.IndexByte([]byte(". -_+"), ch) >= 0 {
				state = 3
			} else {
				state = 0
			}
		case 3:
			if (ch == '\n') || (ch == '\r') {
				outputText = string(verBytes[:idx])
				break
			}
		}
	}
	if len(outputText) == 0 {
		outputText = string(verBytes)
	}
	if (majorVer == 0) && (minorVer == 0) {
		err = ErrUnexpectCommandOutput
	}
	return
}

// Version wrap `go version` command.
func (c *CommandGo) Version() (majorVer, minorVer int, outputText string, err error) {
	out, err := exec.Command(c.exePath(), "version").Output()
	if nil != err {
		return
	}
	return parseVersionBytes(out)
}
