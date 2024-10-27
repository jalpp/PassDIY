package hcpvault

import "strings"

var (
	VAULT_PREFIX           = "hcpvault"
	VAULT_MAIN_DESC        = "Manage Token/Password on Hashicorp Vault"
	VAULT_SUBCOMMAND_NAMES = []string{VAULT_PREFIX + "connect", VAULT_PREFIX + "store", VAULT_PREFIX + "list"}
	VAULT_SUBCOMMAND_DESC  = []string{"Generate HCP API token and connect to Hashicorp Vault", "Store a new secret to Hashicorp Vault", "List HCP Vault secrets log details"}
	VAULT_DISPLAY_COLOR    = "#FFDE59"
)

func ConnectUI() string {
	return Connect()
}

func StoreUI(userInput string) string {
	parts := strings.SplitN(userInput, "=", 2)
	if len(parts) == 2 {
		name := parts[0]
		value := parts[1]
		return Create(name, value)
	}
	return "Invalid format. Use 'name=value'."
}

func ListUI() string {
	var list string = List()
	if strings.Contains(list, "Unauthorized") {
		return "Please connect to Hashicorp vault via hcpvaultconnect"
	}
	return list
}
