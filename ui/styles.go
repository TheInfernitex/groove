package ui

import "github.com/charmbracelet/lipgloss"

var (
	SelectedItemStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("2")). // green
		Bold(true)

	NowPlayingLabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("5")). // magenta
		Bold(true)

	NowPlayingTitleStyle = lipgloss.NewStyle().
		Bold(true)
)
