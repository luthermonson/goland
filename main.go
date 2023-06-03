package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "goland",
		Action:      start,
		Usage:       "",
		Description: "Open goland easily from the cli",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	dir := c.Args().First()

	var err error
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	if strings.Contains(dir, "~") {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		dir = strings.Replace(dir, "~", usr.HomeDir, 1)
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
