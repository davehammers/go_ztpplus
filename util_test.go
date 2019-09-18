package ztp

import (
	"testing"
)

var asciiString = "abcdefghijklmnopqrstuvwxyz0123456789"
var hexString = "61:62:63:64:65:66:67:68:69:6a:6b:6c:6d:6e:6f:70:71:72:73:74:75:76:77:78:79:7a:30:31:32:33:34:35:36:37:38:39"
func TestEncodeHexString(t *testing.T) {
	x := EncodeHexString(asciiString)
	if x != hexString {
		t.Fatal(x)
	}
	t.Log(x)
}
func TestDecodeHexString(t *testing.T) {
	x := DecodeHexString(hexString)
	if x != asciiString {
		t.Fatal(x)
	}
	t.Log(x)
}
