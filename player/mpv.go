package player

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func StartMPV(filename string) *exec.Cmd {
	cmd := exec.Command("mpv", "--input-ipc-server=/tmp/mpvsocket", "--idle", filename)
	cmd.Start()
	// Wait a moment for the socket to be created
	time.Sleep(time.Second) 
	return cmd
}

// SendMPVCommand sends a raw JSON IPC command to the mpv socket.
func SendMPVCommand(command string) {
	socket := "/tmp/mpvsocket"
	cmd := exec.Command("socat", "-", fmt.Sprintf("UNIX-CONNECT:%s", socket))
	cmd.Stdin = strings.NewReader(command)
	
	// We run and wait for socat to complete.
	// We can add error logging here in the future.
	cmd.Run()
}

// PauseMPV now uses the generic SendMPVCommand
func PauseMPV() {
	SendMPVCommand(`{"command": ["cycle", "pause"]}`)
}

// IncreaseVolume increases the MPV volume.
func IncreaseVolume() {
	SendMPVCommand(`{"command": ["add", "volume", "5"]}`)
}

// DecreaseVolume decreases the MPV volume.
func DecreaseVolume() {
	SendMPVCommand(`{"command": ["add", "volume", "-5"]}`)
}
