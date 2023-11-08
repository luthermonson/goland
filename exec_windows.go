package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	root    = "C:" + string(os.PathSeparator)
	basedir = "C:\\Program Files\\JetBrains\\"
	binary  = "\\bin\\goland64.exe"
)

func getExec() string {
	files, err := os.ReadDir(basedir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(), "GoLand") {
			return fmt.Sprintf("%s%s%s", basedir, f.Name(), binary)
		}
	}

	return ""
}

func getArgs(dir string) []string {
	return []string{dir}
}
