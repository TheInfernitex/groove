package ui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/TheInfernitex/groove/player"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Files        []string
	Selected     int
	Playing      bool
	CurrentFile  string
	MPVProcess   *exec.Cmd 
}

func NewModel() Model {
	files := getMP3Files(".")
	return Model{Files: files}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// Kill the MPV process when quitting
			if m.MPVProcess != nil {
				m.MPVProcess.Process.Kill()
			}
			return m, tea.Quit
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
			}
		case "down", "j":
			if m.Selected < len(m.Files)-1 {
				m.Selected++
			}
		case "enter", " ":
			m.Playing = !m.Playing
			m.CurrentFile = m.Files[m.Selected]
			if m.Playing {
				// Start MPV and store the process
				m.MPVProcess = player.StartMPV(m.CurrentFile)
			} else {
				// Pause the MPV process
				player.PauseMPV()
				if m.MPVProcess != nil {
					m.MPVProcess.Process.Kill() 
					m.MPVProcess = nil     
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "🎵 Groove Player\n\nUse ↑/↓ to navigate, [Enter] or [Space] to play/pause\n\n"

	for i, file := range m.Files {
		cursor := "  "
		if m.Selected == i {
			cursor = "👉"
		}
		line := fmt.Sprintf("%s %s\n", cursor, file)
		s += line
	}

	if m.Playing {
		s += fmt.Sprintf("\nNow Playing: %s\n", m.CurrentFile)
	}
	s += "\nPress Q to quit.\n"
	return s
}

func getMP3Files(dir string) []string {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return files
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".mp3") {
			files = append(files, entry.Name())
		}
	}
	return files
}

