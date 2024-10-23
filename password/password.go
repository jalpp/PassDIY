package cmds

import (
	"math/rand"
	"strings"
)

func getLetters() string {
	return "qwertyuiopasdfghjklzxcvbnm"
}

func getSpec() string {
	return "!@#$%^&*()_-+="
}

func getNum() string {
	return "1234567890"
}

func getLetterUpper() string {
	return strings.ToUpper(getLetters())
}

func GetRandomPassword(leng int, pass string) string {

	var buffer string

	var itarray = SplitPassword((pass))

	for i := 0; i < leng; i++ {
		randin := rand.Intn((len(itarray)))
		bufferElement := itarray[randin]
		buffer += bufferElement
	}

	return buffer
}

func CoverUp(pass string) string {
	var buffer string
	for i := 0; i < len(pass); i++ {
		if strings.Compare(SplitPassword(pass)[i], "\n") == 0 {
			buffer += "\n"
		}
		buffer += "*"
	}

	return buffer
}

func SplitPassword(pass string) []string {
	runes := []rune(pass)
	c := make([]string, len(runes))

	for i, r := range runes {
		c[i] = string(r)
	}

	return c
}

func CheckError(input string) bool {
	return strings.Compare(input, "") == 0 || len(input) <= 1 || strings.Contains(strings.ToLower(input), "please")
}

func GetPin() string {
	return GetRandomPassword(PIN_DIGIT_LENGTH, getNum())
}

func GetStrongPassword() string {
	return GetRandomPassword(PASSWORD_CHAR_LENGTH, getLetterUpper()+getLetters()+getNum()+getSpec())
}

func AddSalt(pass string) string {

	if CheckError(pass) {
		return "Salt can not be added! Please generate a password first"
	}

	return GetRandomPassword(rand.Intn(SALT_EXTRA_LENGTH)+2, pass) + pass
}

func GetAPIToken() string {
	return GetRandomPassword(API_TOKEN_CHAR_LENGTH, getLetterUpper()+getLetters())
}
