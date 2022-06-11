package main

import (
	jose "github.com/dvsekhvalnov/jose2go"
)

func main() {
	key := []byte("F3BFC2F8-7512-4063-8461-D209A2E8DA88")
	token, _ := createToken("Obi-Wan Kenobi", "ben@kenobi.com", "Jedi Master", key)
	payload, _, _ := jose.Decode(token, []byte("F3BFC2F8-7512-4063-8461-D209A2E8DA88"))
	println("Signature: " + token)
	println("Verified to: " + payload)
}

func createToken(name string, email string, roles string, key []byte) (string, error) {
	payload := `
{
    "name": , ` + name + `,
    "email": ` + email + `,
    "roles": ` + roles + `
}
`
	return jose.Sign(payload, jose.HS256, key)
}
