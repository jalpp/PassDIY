package cmds

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func HashFunc(pass string) string {

	if CheckError(pass) {
		return "Argon2id hash can not be generated! Please provide token/password input first."
	}

	hash, err := argon2id.CreateHash(pass, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}
