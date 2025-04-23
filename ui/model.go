package ui

import (
	"fmt"

	"github.com/TheInfernitex/groove/player"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Playing bool
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.Playing = !m.Playing
			if m.Playing {
				go player.StartMPV("song.mp3")
			} else {
				go player.PauseMPV()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	status := "Paused"
	if m.Playing {
		status = "Playing"
	}
	return fmt.Sprintf("🎵 Groove Player\n\n[Space] Play/Pause | [Q] Quit\n\nStatus: %s\n", status)
}

