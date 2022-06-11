package main

import (
	jose "github.com/dvsekhvalnov/jose2go"
)

func main() {
	key := []byte("F3BFC2F8-7512-4063-8461-D209A2E8DA88")
	tokenString := formatTokenString("Obi-Wan Kenobi", "ben@kenobi.com", "Jedi Master")
	token, _ := createToken(tokenString, key)
	payload, _, _ := jose.Decode(token, key)
	println("Signature: " + token)
	println("Decoded to: " + payload)
	print("Original String: " + tokenString)
	println("Is the decoded string equal to original string? ", tokenString == payload)
}

func formatTokenString(name string, email string, roles string) string {
	formattedTokenString := `
{
    "name": ` + name + `,
    "email": ` + email + `,
    "roles": ` + roles + `
}
`
	return formattedTokenString
}

func createToken(tokenString string, key []byte) (string, error) {
	return jose.Sign(tokenString, jose.HS256, key)
}
