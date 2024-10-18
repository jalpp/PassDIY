package cmds

import (
	"log"

	"github.com/alexedwards/argon2id"
	"golang.org/x/crypto/bcrypt"
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

func BcryptHash(pass string) string {
	if CheckError(pass) {
		return "BcryptHash hash can not be generated! Please provide token/password input first."
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)

}
