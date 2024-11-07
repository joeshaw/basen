package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/FactomProject/basen"
)

var (
	base36Encoding = basen.NewEncoding("0123456789abcdefghijklmnopqrstuvwxyz")
	base62Encoding = basen.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <from base> <to base> <value>\n", os.Args[0])
		return
	}

	var (
		fromBase = os.Args[1]
		toBase   = os.Args[2]
		value    = os.Args[3]

		encode func([]byte) string
		decode func(string) ([]byte, error)
	)

	switch fromBase {
	case "base36", "b36", "36":
		decode = base36Encoding.DecodeString

	case "base62", "b62", "62":
		decode = base62Encoding.DecodeString

	case "base64", "b64", "64":
		decode = base64.StdEncoding.DecodeString

	case "hex":
		decode = hex.DecodeString

	case "raw":
		decode = func(s string) ([]byte, error) {
			return []byte(s), nil
		}
	}

	switch toBase {
	case "base36", "b36", "36":
		encode = base36Encoding.EncodeToString

	case "base62", "b62", "62":
		encode = base62Encoding.EncodeToString

	case "base64", "b64", "64":
		encode = base64.StdEncoding.EncodeToString

	case "hex":
		encode = hex.EncodeToString

	case "raw":
		encode = func(b []byte) string {
			return string(b)
		}
	}

	decoded, err := decode(value)
	if err != nil {
		fmt.Println("Error decoding value:", err)
		return
	}

	fmt.Println(encode(decoded))
}
