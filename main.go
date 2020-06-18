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
		Name:   "goland",
		Action: start,
		Usage:  "",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	executable := getExecutable()
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
			log.Fatal(err)
		}
		dir = strings.Replace(dir, "~", usr.HomeDir, 1)
	}

	if executable == "" || dir == "" {
		return errors.New("empty executable or directory")
	}

	cmd := exec.Command(executable, dir)
	cmd.Start()

	return nil
}
