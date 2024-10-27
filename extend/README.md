## Config custom vault to use PassDIY TUI

to config custom vaults that are not currently supported by Passdiy all you have to do is edit the interface.go file and define your custom implementation of the functions, then you set `export USE_PASDIY_CUSTOM_VAULT=true` and PassDIY will automatically interface the custom vault

```go
package extend

var (
	VAULT_PREFIX           = "pref"
	VAULT_MAIN_DESC        = "Manage token/password on " + VAULT_PREFIX
	VAULT_SUBCOMMAND_NAMES = []string{VAULT_PREFIX + "store", VAULT_PREFIX + "list"}
	VAULT_SUBCOMMAND_DESC  = []string{"store", "lists"}
	VAULT_DISPLAY_COLOR    = "#E2EAF4"
)

func ConnectUI() string {
	return Connect()
}

func StoreUI(userInput string) string {

	var parser string

	return Create(userInput, parser)
}

func ListUI() string {
	return List()
}

```