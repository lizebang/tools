package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	help = `%s usage:
%s [ -h | -H | --help ]
%s [path to store file]
`
)

func getPath() string {
	args := os.Args
	if len(args) == 1 {
		path, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		return path
	} else if len(args) == 2 {
		if isHelp(args[1]) {
			fmt.Fprintf(os.Stdout, help, args[0], args[0], args[0])
			os.Exit(0)
		}
		path, err := filepath.Abs(args[1])
		if err != nil {
			panic(err)
		}
		os.MkdirAll(path, 0755)
		return path
	}
	panic(errors.New("too many arguments"))
}

func isHelp(s string) bool {
	if s == "-h" || s == "-H" || s == "--help" {
		return true
	}
	return false
}
