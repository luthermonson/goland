// +build windows

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	basedir = "C:\\Program Files\\JetBrains\\"
	binary  = "\\bin\\goland64.exe"
)

func getExecutable() string {
	files, err := ioutil.ReadDir(basedir)
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
