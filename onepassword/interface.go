package onepassword

import "strings"

func StoreUI(userInput string) string {
	parts := strings.SplitN(userInput, "|", 3)
	if len(parts) == 3 {
		user := parts[0]
		pass := parts[1]
		url := parts[2]
		return Create(user, pass, url)
	}
	return "Invalid format. use 'user|value|url'."
}

func ListUI() string {
	return List()
}
