package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	GreenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7DDA58"))
	ErrorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#E4080A"))
	VaultStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFDE59"))
	OPassStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#4CAAF7"))
	ConfigStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#CC6CE7"))
)
