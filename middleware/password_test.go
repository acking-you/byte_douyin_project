package middleware

import "testing"

//40bd001563085fc35165329ea1ff5c5ecbdbbeef
func TestSHA1(t *testing.T) {
	print(SHA1("123"))
}
