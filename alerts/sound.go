package alerts

import (
	"fmt"
	"log"
	"os/exec"
)

func Play() {
	sound := fmt.Sprintf("%s/data/alarm.oga", curdir())

	cmd := exec.Command("/usr/bin/cvlc", "--play-and-exit", "--no-loop", sound)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	//log.Printf("Command finished with error: %v", err)
}
