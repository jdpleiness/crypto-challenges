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
		fmt.Println("Error, string lengths do not match or key is not a single byte")
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

	for i := range rawslice {
		rawslice[i] = rawinput[i] ^ rawkey[i]
	}

	return hex.EncodeToString(rawslice)
}

func singleCharXor(input string, key byte) (out []byte) {
	rawinput, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	rawslice := make([]byte, len(rawinput))

	for i := range rawslice {
		rawslice[i] = rawinput[i] ^ key
	}

	return rawslice
}

func freqScore(input []byte) int {
	/*
		Score based on how many ascii chars it recognizes;
		one point for normal, two for e/E, and three for space
	*/
	freqscore := 0

	for _, char := range input {
		if char > 31 && char < 127 {
			freqscore++

			if char == 32 {
				freqscore += 2
			} else if char == 69 || char == 101 {
				freqscore++
			}
		}
	}

	return freqscore
}

func singleByteXorDecode(input string) (outkey string, outmsg string) {
	highscore := 0
	key := byte(0)

	for char := byte(33); char <= 126; char++ {
		score := freqScore(singleCharXor(input, char))

		if score > highscore {
			highscore = score
			key = char
		}
	}

	return string(key), string(singleCharXor(input, key))
}
