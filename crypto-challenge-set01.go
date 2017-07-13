package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
}

func hexToBase64(src string) (out string) {
	data, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	return base64.StdEncoding.EncodeToString(data)
}

func xor(input string, key string) (out string) {
	// function only works when input and key are same length
	if len(input) != len(key) {
		fmt.Println("Error, string lengths do not match")
		return
	}

	rawinput, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	rawkey, err := hex.DecodeString(key)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	rawslice := make([]byte, len(rawinput))

	for i := 0; i < len(rawslice); i++ {
		rawslice[i] = rawinput[i] ^ rawkey[i]
	}

	return hex.EncodeToString(rawslice)
}

func singleByteXorDecode(input string) (outkey string, outmsg string) {
	rawinput, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	//TODO learn more about scoring character freq.

}
