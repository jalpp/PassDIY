package onepassword

import "strings"

var (
	VAULT_PREFIX           = "1pass"
	VAULT_MAIN_DESC        = "Manage Token/Password on 1Password"
	VAULT_SUBCOMMAND_NAMES = []string{VAULT_PREFIX + "store", VAULT_PREFIX + "list"}
	VAULT_SUBCOMMAND_DESC  = []string{"Store a new secret to 1Password in password format", "List 1Password Vault item names"}
	VAULT_DISPLAY_COLOR    = "#4CAAF7"
)

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
