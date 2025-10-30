# 🎵 groove

[![Hacktoberfest](https://img.shields.io/badge/Hacktoberfest-Participating-2b9348?style=flat-square)](https://hacktoberfest.com/)

A terminal-based music player built with Go, [Bubbletea](https://github.com/charmbracelet/bubbletea), [MPV](https://mpv.io/), and [SoX](http://sox.sourceforge.net/).

---

## ✨ Features

- Scans the current directory for `.mp3` files.
- Interactive TUI to select, play, and pause/unpause songs.
- Real-time volume control.

## ⌨️ Controls

| Key               | Action                 |
| :---------------- | :--------------------- |
| `↑` / `k`         | Move selection up      |
| `↓` / `j`         | Move selection down    |
| `Enter` / `Space` | Play / Pause / Unpause |
| `+` / `=`         | Increase Volume        |
| `-`               | Decrease Volume        |
| `q` / `Ctrl+c`    | Quit                   |

## ⚙️ Requirements

- Go (>= 1.20)
- MPV media player
- socat (for IPC communication with MPV)
- SoX (optional, for audio effects)

## Getting Started

```bash
git clone [https://github.com/TheInfernitex/groove.git](https://github.com/TheInfernitex/groove.git)
cd groove
go mod tidy
go run .
```

## 📌 Upcoming Features

- Recursive file browser to choose music (not just root)
- Visual song progress bar
- Visual volume slider
- Lipgloss UI styling for a polished look
- SoX integration for effects (Pitch, Speed, etc.)

## 🤝 Contributions

This project is actively participating in **Hacktoberfest**\!

We welcome all contributions. Feel free to fork, star, and PR improvements\! Please check out the [**CONTRIBUTING.md**](https://github.com/TheInfernitex/groove/blob/main/CONTRIBUTING.md) file for guidelines on how to make a valid contribution.

## Author

Built by Parth Agarwal with love and a kitty terminal.
