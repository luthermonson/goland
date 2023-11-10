package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error starting goland: %s", err)
	}
}

func run() error {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if dir == "" {
		dir = cwd
	}

	if strings.Contains(dir, "~") {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		dir = strings.Replace(dir, "~", usr.HomeDir, 1)
	}

	if !strings.HasPrefix(dir, root) {
		dir = fmt.Sprintf("%s%s%s", cwd, string(os.PathSeparator), dir)
	}
	if dir == "" {
		return errors.New("directory to open is empty, try again")
	}

	name := getExec()
	if name == "" {
		return errors.New("unable to find goland executable")
	}

	return exec.Command(name, getArgs(dir)...).Start()
}
