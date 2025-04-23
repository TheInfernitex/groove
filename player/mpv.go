package player

import (
	"fmt"
	"os/exec"
	"time"
)

func StartMPV(filename string) *exec.Cmd {
	cmd := exec.Command("mpv", "--input-ipc-server=/tmp/mpvsocket", filename)
	cmd.Start()
	time.Sleep(time.Second)
	return cmd
}

func PauseMPV() {
	socket := "/tmp/mpvsocket"
	cmd := exec.Command("echo", `{"command": ["cycle", "pause"]}`)
	mpv := exec.Command("socat", "-", fmt.Sprintf("UNIX-CONNECT:%s", socket))
	pipe, _ := cmd.StdoutPipe()
	cmd.Start()
	mpv.Stdin = pipe
	mpv.Start()
	cmd.Wait()
	mpv.Wait()
}

