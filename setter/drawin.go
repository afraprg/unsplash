package setter

import (
	"os/exec"
	"strconv"
)

func SetFromFile(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}
