package cmds

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	hcp "github.com/jalpp/passdiy/hcpvault"
	cmd "github.com/jalpp/passdiy/password"
)

type CommandItem struct {
	title, desc string
}

const (
	passDesc            = "strong password"
	tokenDesc           = "strong API token"
	pinDesc             = "strong 6-digit pin"
	pwpDesc             = "strong passphrase"
	saltDesc            = "password with extra salt on top"
	hashDesc            = "hash value of a password with Argon2id"
	directDesc          = "direct hash running buffer value with Argon2id"
	hcpvaultstoreDesc   = "Store a new secret to Hashicop Vault"
	hcpvaultconnectDesc = "Generate HCP API token and connect to Hashicop Vault"
)

func (i CommandItem) Title() string       { return i.title }
func (i CommandItem) Description() string { return i.desc }
func (i CommandItem) FilterValue() string { return i.title }

func GetSingleCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s", cmd)
}

func GetMulCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate 5 multiple %s all at once", cmd)
}

func GetHundCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s by generating 100s random %s and randomly picking one", cmd, cmd)
}

func GetTenKCommandInfo(cmd string) string {
	return fmt.Sprintf("Generate a single %s by generating 10000s random %s and randomly picking one", cmd, cmd)
}

func CreateCommandItems() []list.Item {
	return []list.Item{
		CommandItem{title: "pass", desc: GetSingleCommandInfo(passDesc)},
		CommandItem{title: "passmmul", desc: GetMulCommandInfo(passDesc)},
		CommandItem{title: "pass100", desc: GetHundCommandInfo(passDesc)},
		CommandItem{title: "pass10000", desc: GetTenKCommandInfo(passDesc)},
		CommandItem{title: "token", desc: GetSingleCommandInfo(tokenDesc)},
		CommandItem{title: "tokenmul", desc: GetMulCommandInfo(tokenDesc)},
		CommandItem{title: "token100", desc: GetHundCommandInfo(tokenDesc)},
		CommandItem{title: "token10000", desc: GetTenKCommandInfo(tokenDesc)},
		CommandItem{title: "pin", desc: GetSingleCommandInfo(pinDesc)},
		CommandItem{title: "pinmul", desc: GetMulCommandInfo(pinDesc)},
		CommandItem{title: "pin100", desc: GetHundCommandInfo(pinDesc)},
		CommandItem{title: "pin10000", desc: GetTenKCommandInfo(pinDesc)},
		CommandItem{title: "pwp", desc: GetSingleCommandInfo(pwpDesc)},
		CommandItem{title: "salt", desc: GetSingleCommandInfo(saltDesc)},
		CommandItem{title: "hash", desc: GetSingleCommandInfo(hashDesc)},
		CommandItem{title: "directhash", desc: GetSingleCommandInfo(directDesc)},
		CommandItem{title: "hcpvaultstore", desc: hcpvaultstoreDesc},
		CommandItem{title: "hcpvaultconnect", desc: hcpvaultconnectDesc},
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
	switch strings.TrimSpace(input) {
	case "pass":
		return cmd.GetStrongPassword()
	case "passmul":
		return cmd.GetMul("pass")
	case "pass100":
		return cmd.GetHundPick("pass")
	case "pass10000":
		return cmd.GetTenKPick("pass")
	case "token":
		return cmd.GetAPIToken()
	case "tokenmul":
		return cmd.GetMul("token")
	case "token100":
		return cmd.GetHundPick("token")
	case "token10000":
		return cmd.GetTenKPick("token")
	case "pin":
		return cmd.GetPin()
	case "pinmul":
		return cmd.GetMul("pin")
	case "pin100":
		return cmd.GetHundPick("pin")
	case "pin10000":
		return cmd.GetTenKPick("pin")
	case "pwp":
		return cmd.GetPwp()
	case "salt":
		return cmd.AddSalt(userInput)
	case "hash":
		return cmd.HashFunc(userInput)
	case "hcpvaultstore":
		parts := strings.SplitN(userInput, "=", 2)
		if len(parts) == 2 {
			name := parts[0]
			value := parts[1]
			return hcp.Create(name, value)
		}
		return "Invalid format. Use 'name=value'."
	case "directhash":
		return cmd.HashFunc(userInput)
	case "hcpvaultconnect":
		return hcp.Connect()
	default:
		return fmt.Sprintf("Unknown command: %s", input)
	}
}

func CoverUp(pass string) string {
	return cmd.CoverUp(pass)
}
