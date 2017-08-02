package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if hexToBase64(input) != expected {
		t.Errorf("string returned was not expected")
	}
}

func TestXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	if xor(input, key) != expected {
		t.Errorf("string returned was not expected")
	}

}

func TestSingleCharXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	key := byte(33)
	expected := "3d2030213e202021273b236a72727128393d"

	if hex.EncodeToString(singleCharXor(input, key)) != (expected) {
		fmt.Println(hex.EncodeToString(singleCharXor(input, key)))
		t.Errorf("string returned was not expected")
	}
}

func TestFreqScore(t *testing.T) {
	input := "Hello World!"
	expected := 15
	if freqScore([]byte(input)) != expected {
		t.Errorf("Did not return expected int value")
	}
}

func TestSingleByteXorDecode(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedmsg := "Cooking MC's like a pound of bacon"
	expcetedkey := "X"

	key, msg := singleByteXorDecode(input)

	if msg != expectedmsg || key != expcetedkey {
		fmt.Println("message: ", msg)
		fmt.Println("key: ", key)
		t.Errorf("key or msg did not match expected")
	}

}
