package cmds

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	custom "github.com/jalpp/passdiy/extend"
	hcp "github.com/jalpp/passdiy/hcpvault"
	opass "github.com/jalpp/passdiy/onepassword"
	cmd "github.com/jalpp/passdiy/password"
)

type CommandItem struct {
	title, desc string
	Subcmd      []CommandItem
}

var (
	pwpDesc             = fmt.Sprintf("Generate strong %d-word passphrase", cmd.PASSPHRASE_COUNT_NUM)
	saltDesc            = fmt.Sprintf("Generate password with extra %d-char salt on top", cmd.SALT_EXTRA_LENGTH)
	hashDesc            = "Generate hash value of a password with Argon2id or bcrypthash"
	argonhashDesc       = "Generate hash value of a password with Argon2id"
	bcrpthashDesc       = "Generate hash value of a password with bcrypt algorithm"
	hcpvaultstoreDesc   = hcp.VAULT_SUBCOMMAND_DESC[1]
	hcpvaultconnectDesc = hcp.VAULT_SUBCOMMAND_DESC[0]
	hcpvaultlistDesc    = hcp.VAULT_SUBCOMMAND_DESC[2]
	opassstoreDesc      = opass.VAULT_SUBCOMMAND_DESC[0]
	opasslistDesc       = opass.VAULT_SUBCOMMAND_DESC[1]
	mainpassDesc        = "Generate strong passwords from various algorithms"
	mainpinDesc         = "Generate strong pins from various algorithms"
	maintokenDesc       = "Generate strong token from various algorithms"
	hcpDesc             = hcp.VAULT_MAIN_DESC
	opassDesc           = opass.VAULT_MAIN_DESC
	configDesc          = "Config PassDIY password, token, pin, salt char lengths"
)

const cf = cmd.LOTTERY_WHEEL_COUNT

var (
	configItems = []CommandItem{
		{title: "configpass", desc: cmd.CONFIG_PASS_DESC},
		{title: "configtoken", desc: cmd.CONFIG_TOKEN_DESC},
		{title: "configpin", desc: cmd.CONFIG_PIN_DESC},
		{title: "configpwp", desc: cmd.CONFIG_PWP_WORD_DESC},
		{title: "configmul", desc: cmd.CONFIG_MULTI_DESC},
		{title: "configsalt", desc: cmd.CONFIG_SALT_DESC},
	}
)

func (i CommandItem) Title() string       { return i.title }
func (i CommandItem) Description() string { return i.desc }
func (i CommandItem) FilterValue() string { return i.title }

func GetSingleCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s", cmd)
}

func GetMulCommandInfo(cmds string) string {
	return fmt.Sprintf("Generate %d multiple %s all at once", cmd.MULTIPLE_VALUE_COUNT, cmds)
}

func GetHundCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s by generating 100s and randomly picking one", cmd)
}

func GetTenKCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s by generating 10000s and randomly picking one", cmd)
}

func CreateCommandItems() []list.Item {

	passItems := []CommandItem{
		{title: "pass", desc: GetSingleCommandInfo("strong password")},
		{title: "passmul", desc: GetMulCommandInfo("password")},
		{title: fmt.Sprintf("pass%d", cf), desc: GetHundCommandInfo("password")},
		{title: fmt.Sprintf("pass%d", cf*cf), desc: GetTenKCommandInfo("password")},
	}

	pinItems := []CommandItem{
		{title: "pin", desc: GetSingleCommandInfo("pin")},
		{title: "pinmul", desc: GetMulCommandInfo("pin")},
		{title: fmt.Sprintf("pin%d", cf), desc: GetHundCommandInfo("pin")},
		{title: fmt.Sprintf("pin%d", cf*cf), desc: GetTenKCommandInfo("pin")},
	}

	tokenItems := []CommandItem{
		{title: "token", desc: GetSingleCommandInfo("token")},
		{title: "tokenmul", desc: GetMulCommandInfo("token")},
		{title: fmt.Sprintf("token%d", cf), desc: GetHundCommandInfo("token")},
		{title: fmt.Sprintf("token%d", cf*cf), desc: GetTenKCommandInfo("token")},
	}
	hashItems := []CommandItem{
		{title: "argonhash", desc: argonhashDesc},
		{title: "bcrypthash", desc: bcrpthashDesc},
	}

	hcpItems := []CommandItem{
		{title: hcp.VAULT_SUBCOMMAND_NAMES[1], desc: hcpvaultstoreDesc},
		{title: hcp.VAULT_SUBCOMMAND_NAMES[0], desc: hcpvaultconnectDesc},
		{title: hcp.VAULT_SUBCOMMAND_NAMES[2], desc: hcpvaultlistDesc},
	}

	opassItems := []CommandItem{
		{title: opass.VAULT_SUBCOMMAND_NAMES[0], desc: opassstoreDesc},
		{title: opass.VAULT_SUBCOMMAND_NAMES[1], desc: opasslistDesc},
	}

	customItems := []CommandItem{
		{title: custom.VAULT_SUBCOMMAND_NAMES[0], desc: custom.VAULT_SUBCOMMAND_DESC[0]},
		{title: custom.VAULT_SUBCOMMAND_NAMES[1], desc: custom.VAULT_SUBCOMMAND_DESC[1]},
	}

	if strings.ToLower(os.Getenv("USE_PASDIY_CUSTOM_VAULT")) == "true" {
		return []list.Item{
			CommandItem{title: "pass", desc: mainpassDesc, Subcmd: passItems},
			CommandItem{title: "pin", desc: mainpinDesc, Subcmd: pinItems},
			CommandItem{title: "token", desc: maintokenDesc, Subcmd: tokenItems},
			CommandItem{title: "salt", desc: saltDesc},
			CommandItem{title: "pwp", desc: pwpDesc},
			CommandItem{title: "config", desc: configDesc, Subcmd: configItems},
			CommandItem{title: "hash", desc: hashDesc, Subcmd: hashItems},
			CommandItem{title: hcp.VAULT_PREFIX, desc: hcpDesc, Subcmd: hcpItems},
			CommandItem{title: opass.VAULT_PREFIX, desc: opassDesc, Subcmd: opassItems},
			CommandItem{title: custom.VAULT_PREFIX, desc: custom.VAULT_MAIN_DESC, Subcmd: customItems},
		}
	}

	return []list.Item{
		CommandItem{title: "pass", desc: mainpassDesc, Subcmd: passItems},
		CommandItem{title: "pin", desc: mainpinDesc, Subcmd: pinItems},
		CommandItem{title: "token", desc: maintokenDesc, Subcmd: tokenItems},
		CommandItem{title: "salt", desc: saltDesc},
		CommandItem{title: "pwp", desc: pwpDesc},
		CommandItem{title: "config", desc: configDesc, Subcmd: configItems},
		CommandItem{title: "hash", desc: hashDesc, Subcmd: hashItems},
		CommandItem{title: hcp.VAULT_PREFIX, desc: hcpDesc, Subcmd: hcpItems},
		CommandItem{title: opass.VAULT_PREFIX, desc: opassDesc, Subcmd: opassItems},
	}
}

type CommandFinishedMsg struct {
	result string
}

func (i CommandFinishedMsg) Result() string { return i.result }

func ExecuteCommand(command, input string) tea.Cmd {
	return func() tea.Msg {
		result := HandleCommand(command, input)
		time.Sleep(1 * time.Second)
		return CommandFinishedMsg{result: result}
	}
}

func HandleCommand(input, userInput string) string {
	const config = cmd.LOTTERY_WHEEL_COUNT
	switch strings.TrimSpace(input) {
	case "pass":
		return cmd.GetStrongPassword()
	case "passmul":
		return cmd.GetMul("pass")
	case fmt.Sprintf("pass%d", config):
		return cmd.GetHundPick("pass")
	case fmt.Sprintf("pass%d", (config * config)):
		return cmd.GetTenKPick("pass")
	case "token":
		return cmd.GetAPIToken()
	case "tokenmul":
		return cmd.GetMul("token")
	case fmt.Sprintf("token%d", config):
		return cmd.GetHundPick("token")
	case fmt.Sprintf("token%d", config*config):
		return cmd.GetTenKPick("token")
	case "pin":
		return cmd.GetPin()
	case "pinmul":
		return cmd.GetMul("pin")
	case fmt.Sprintf("pin%d", config):
		return cmd.GetHundPick("pin")
	case fmt.Sprintf("pin%d", config*config):
		return cmd.GetTenKPick("pin")
	case "pwp":
		return cmd.GetPwp()
	case "salt":
		return cmd.AddSalt(userInput)
	case "argonhash":
		return cmd.HashFunc(userInput)
	case "bcrypthash":
		return cmd.BcryptHash(userInput)
	case "configpass":
		return cmd.SetPasswordLength(userInput)
	case "configtoken":
		return cmd.SetAPITokenLength(userInput)
	case "configpin":
		return cmd.SetPinLength(userInput)
	case "configpwp":
		return cmd.SetPwpWordCount(userInput)
	case "configmul":
		return cmd.SetMulCount(userInput)
	case "configsalt":
		return cmd.SetSaltLength(userInput)
	case hcp.VAULT_SUBCOMMAND_NAMES[1]:
		return hcp.StoreUI(userInput)
	case hcp.VAULT_SUBCOMMAND_NAMES[0]:
		return hcp.ConnectUI()
	case hcp.VAULT_SUBCOMMAND_NAMES[2]:
		return hcp.ListUI()
	case opass.VAULT_SUBCOMMAND_NAMES[0]:
		return opass.StoreUI(userInput)
	case opass.VAULT_SUBCOMMAND_NAMES[1]:
		return opass.ListUI()
	case custom.VAULT_SUBCOMMAND_NAMES[0]:
		return custom.StoreUI(userInput)
	case custom.VAULT_SUBCOMMAND_NAMES[1]:
		return custom.ListUI()
	default:
		return fmt.Sprintf("Unknown command: %s", input)
	}
}

func CoverUp(pass string) string {
	return cmd.CoverUp(pass)
}

func IsConfigCommand(command string) bool {
	return strings.Contains(command, "config")
}

func IsHashCommand(command string) bool {
	return command == "bcrypthash" || command == "argonhash"
}

func IsCommandInputMode(commandName string) bool {
	return IsConfigCommand(commandName) || IsHashCommand(commandName) || commandName == "hcpvaultstore" || commandName == "1passstore"
}
