package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abcdfgfdg343%$%$&#@!~"

	// default encoding/decoding for base64

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	// URL encoding/decoding for base64

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
