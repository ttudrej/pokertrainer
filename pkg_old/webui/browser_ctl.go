package webui

import (
	"os/exec"
	"runtime"

	"github.com/ttudrej/pokertrainer/debugging"
)

// #################################################################
// open opens the specified URL in the default browser of the user.
func myOpenURI(url string) error {
	Info.Println(debugging.ThisFunc())
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "/Applications/Firefox.app/Contents/MacOS/firefox"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	// return

	xErr := exec.Command(cmd, args...).Start()
	Info.Println("Start err: ", xErr)
	return xErr
	// return exec.Command(cmd, args...).Run()

}

// #################################################################
