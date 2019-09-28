package rungo

import (
	"testing"
)

func doTestParseVersionBytesSuccess(t *testing.T, testText, expOutText string, expMajorVer, expMinorVer int) {
	if majVer, minVer, outText, err := parseVersionBytes([]byte(testText)); nil != err {
		t.Errorf("parsing failed: %v", err)
	} else if outText != expOutText {
		t.Errorf("unexpect output text (expect %v): %v", expOutText, outText)
	} else if (majVer != expMajorVer) || (minVer != expMinorVer) {
		t.Errorf("unexpect version (expect %d, %d): %d, %d", expMajorVer, expMinorVer, majVer, minVer)
	}
}

func TestParseVersionBytes(t *testing.T) {
	doTestParseVersionBytesSuccess(t,
		"go version go1.13 darwin/amd64\n",
		"go version go1.13 darwin/amd64", 1, 13)
	doTestParseVersionBytesSuccess(t,
		"go version go1.12.5 darwin/amd64\n",
		"go version go1.12.5 darwin/amd64", 1, 12)
}
