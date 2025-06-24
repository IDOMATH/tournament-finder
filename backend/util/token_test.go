package util

import (
	"fmt"
	"testing"
)

func TestTokenId(t *testing.T) {
	expected := 17
	token := MakeToken(expected)
	got := GetUserIdFromToken(token)
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}

	expected = 983
	token = MakeToken(expected)
	got = GetUserIdFromToken(token)
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}

	expected = 2983
	token = MakeToken(expected)
	fmt.Println(token)
	got = GetUserIdFromToken(token)
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}

	expected = 629834
	token = MakeToken(expected)
	fmt.Println(token)
	got = GetUserIdFromToken(token)
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}

	expected = 999999
	token = MakeToken(expected)
	fmt.Println(token)
	got = GetUserIdFromToken(token)
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}
}
