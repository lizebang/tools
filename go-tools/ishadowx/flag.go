package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	help = `%s usage:
%s [ -h | -H | --help | --Help ]  how to use %s
%s [ The path to save the file ]  defualt path is "."
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
	} else if len(args) > 2 {
		fmt.Fprintf(os.Stderr, "Error: too many arguments\n"+help, args[0], args[0], args[0], args[0])
		os.Exit(1)
	}
	if isHelp(args[1]) {
		fmt.Fprintf(os.Stderr, help, args[0], args[0], args[0], args[0])
		os.Exit(1)
	}
	path, err := filepath.Abs(args[1])
	if err != nil {
		panic(err)
	}
	os.MkdirAll(path, 0755)
	return path
}

func isHelp(s string) bool {
	if s == "-h" || s == "-H" || s == "--help" || s == "--Help" {
		return true
	}
	return false
}
