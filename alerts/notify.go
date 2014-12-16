package alerts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Notify(message string) {
	dir := curdir()
	icon := fmt.Sprintf("%s/data/tomato.png", dir)

	cmd := exec.Command("/usr/bin/notify-send", "-i", icon, message)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func curdir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
