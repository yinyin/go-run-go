package rungo

import (
	"testing"
)

func TestConvertDigitByte(t *testing.T) {
	if digit, ok := convertDigitByte('0'); !ok || (digit != 0) {
		t.Errorf("expect digit = 0, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('1'); !ok || (digit != 1) {
		t.Errorf("expect digit = 1, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('2'); !ok || (digit != 2) {
		t.Errorf("expect digit = 2, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('3'); !ok || (digit != 3) {
		t.Errorf("expect digit = 3, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('8'); !ok || (digit != 8) {
		t.Errorf("expect digit = 8, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('9'); !ok || (digit != 9) {
		t.Errorf("expect digit = 9, ok = true (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('a'); ok {
		t.Errorf("expect ok = false (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte('.'); ok {
		t.Errorf("expect ok = false (got digit = %d, ok = %v)", digit, ok)
	}
	if digit, ok := convertDigitByte(' '); ok {
		t.Errorf("expect ok = false (got digit = %d, ok = %v)", digit, ok)
	}
}
