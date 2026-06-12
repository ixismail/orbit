package theme

import "github.com/charmbracelet/lipgloss"

const (
	BrightGreen = lipgloss.Color("46")
	Cyan        = lipgloss.Color("86")
	Red         = lipgloss.Color("196")
	Orange      = lipgloss.Color("214")
	Pink        = lipgloss.Color("205")
	Black       = lipgloss.Color("0")
)

// Theme defines the global color palette and typography for Orbit
var (
	Success = lipgloss.NewStyle().Foreground(BrightGreen).Bold(true)
	Info    = lipgloss.NewStyle().Foreground(Cyan)
	Error   = lipgloss.NewStyle().Foreground(Red).Bold(true)
	Warning = lipgloss.NewStyle().Foreground(Orange)
)