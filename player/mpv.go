package player

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func StartMPV(filename string) {
	exec.Command("mpv", "--input-ipc-server=/tmp/mpvsocket", filename).Start()
	time.Sleep(time.Second)
}

func PauseMPV() {
	socket := "/tmp/mpvsocket"
	if _, err := os.Stat(socket); os.IsNotExist(err) {
		return
	}
	cmd := exec.Command("echo", `{"command": ["cycle", "pause"]}`)
	mpv := exec.Command("socat", "-", fmt.Sprintf("UNIX-CONNECT:%s", socket))
	pipe, _ := cmd.StdoutPipe()
	cmd.Start()
	mpv.Stdin = pipe
	mpv.Start()
	cmd.Wait()
	mpv.Wait()
}

