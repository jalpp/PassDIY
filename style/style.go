package style

import (
	"github.com/charmbracelet/lipgloss"
	custom "github.com/jalpp/passdiy/extend"
	hcp "github.com/jalpp/passdiy/hcpvault"
	opass "github.com/jalpp/passdiy/onepassword"
)

var (
	GreenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7DDA58"))
	ErrorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#E4080A"))
	VaultStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(hcp.VAULT_DISPLAY_COLOR))
	OPassStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(opass.VAULT_DISPLAY_COLOR))
	CustomStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(custom.VAULT_DISPLAY_COLOR))
	ConfigStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#CC6CE7"))
)
