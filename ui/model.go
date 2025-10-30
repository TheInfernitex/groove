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
			selectedFile := m.Files[m.Selected]

			// Case 1: No process is active OR a different song is selected
			if m.MPVProcess == nil || m.CurrentFile != selectedFile {
				// Kill old process if it exists (for the "different song" case)
				if m.MPVProcess != nil {
					m.MPVProcess.Process.Kill()
				}
				// Start new song
				m.CurrentFile = selectedFile
				m.MPVProcess = player.StartMPV(m.CurrentFile)
				m.Playing = true
			} else {
				// Case 2: A process for this song is already active, so just toggle pause
				player.PauseMPV()
				m.Playing = !m.Playing // Toggle the playing state
			}
		case "+", "=": // Use "=" as well, since "+" often requires Shift
			player.IncreaseVolume()
		case "-":
			player.DecreaseVolume()
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "🎵 Groove Player\n\n"
	s += "Use ↑/↓ to navigate\n"
	s += "[Enter] or [Space] to play/pause\n"
	s += "[+/-] to change volume\n\n"

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

