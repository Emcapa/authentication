package main

import (
	jose "github.com/dvsekhvalnov/jose2go"
	"testing"
)

func TestCreate(t *testing.T) {
	key := []byte("F3BFC2F8-7512-4063-8461-D209A2E8DA88")
	tokenString := formatTokenString("Obi-Wan Kenobi", "ben@kenobi.com", "Jedi Master")
	token, tokenErr := createToken(tokenString, key)
	payload, _, decodeErr := jose.Decode(token, key)
	if tokenErr != nil {
		t.Errorf("The jose library encountered an error while creating the token")
	} else if decodeErr != nil {
		t.Errorf("The jose library encountered an error while decoding the token")
	} else if tokenString != payload {
		t.Errorf("The original token doesn't match the decoded token")
	}
}
