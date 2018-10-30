package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	apiContainer      = "api"
	commentsContainer = "comments"
	reviewContainer   = "reviews"
	dumpContainer     = "dump"
)

func restartContainer(cName string) error {
	_, err := exec.Command("docker", "restart", cName).Output()
	if err != nil {
		return err
	}
	return nil

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("No arguments")
	}

	switch args[1] {
	case apiContainer:
		handleError(restartContainer("reviewzone_api"))

	case commentsContainer:

		handleError((restartContainer("reviewzone_comments")))
		time.Sleep(1 * time.Second)
		handleError(restartContainer("reviewzone_api"))

		fmt.Println("containers restarted.")

	case reviewContainer:

		handleError(restartContainer("reviewzone_reviewer"))
		time.Sleep(1 * time.Second)
		handleError(restartContainer("reviewzone_api"))

	case dumpContainer:

		handleError(restartContainer("reviewzone_dump"))
		time.Sleep(1 * time.Second)
		handleError(restartContainer("reviewzone_api"))

	default:
		fmt.Println("Invalid container.")
	}

}
