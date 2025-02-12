package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/FactomProject/basen"
)

var (
	base36Encoding = basen.NewEncoding("0123456789abcdefghijklmnopqrstuvwxyz")
	base62Encoding = basen.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <from base> <to base> <value>\n", os.Args[0])
		fmt.Println()
		fmt.Println("Available bases:")
		fmt.Println("  - base36, b36, 36")
		fmt.Println("  - base62, b62, 62")
		fmt.Println("  - base64, b64, 64")
		fmt.Println("  - base10, b10, 10, int")
		fmt.Println("  - hex")
		fmt.Println("  - raw")
		fmt.Println("  - custom:<alphabet>")
		return
	}

	var (
		fromBase = os.Args[1]
		toBase   = os.Args[2]
		value    = os.Args[3]

		encode func([]byte) string
		decode func(string) ([]byte, error)
	)

	fromBase, fromAlpha, _ := strings.Cut(fromBase, ":")
	toBase, toAlpha, _ := strings.Cut(toBase, ":")

	switch fromBase {
	case "base36", "b36", "36":
		decode = base36Encoding.DecodeString

	case "base62", "b62", "62":
		decode = base62Encoding.DecodeString

	case "base64", "b64", "64":
		decode = base64.StdEncoding.DecodeString

	case "base10", "b10", "10", "int":
		decode = func(s string) ([]byte, error) {
			var v big.Int
			if _, ok := v.SetString(s, 10); !ok {
				return nil, fmt.Errorf("invalid int value")
			}
			return v.Bytes(), nil
		}

	case "hex":
		decode = hex.DecodeString

	case "raw":
		decode = func(s string) ([]byte, error) {
			return []byte(s), nil
		}

	case "custom":
		enc := basen.NewEncoding(fromAlpha)
		decode = enc.DecodeString

	default:
		fmt.Println("Unknown \"from\" base:", fromBase)
		return
	}

	switch toBase {
	case "base36", "b36", "36":
		encode = base36Encoding.EncodeToString

	case "base62", "b62", "62":
		encode = base62Encoding.EncodeToString

	case "base64", "b64", "64":
		encode = base64.StdEncoding.EncodeToString

	case "base10", "b10", "10", "int":
		encode = func(b []byte) string {
			var v big.Int
			v.SetBytes(b)
			return v.String()
		}

	case "hex":
		encode = hex.EncodeToString

	case "raw":
		encode = func(b []byte) string {
			return string(b)
		}

	case "custom":
		enc := basen.NewEncoding(toAlpha)
		encode = enc.EncodeToString

	default:
		fmt.Println("Unknown \"to\" base:", toBase)
		return
	}

	decoded, err := decode(value)
	if err != nil {
		fmt.Println("Error decoding value:", err)
		return
	}

	fmt.Println(encode(decoded))
}
