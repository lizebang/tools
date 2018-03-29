package main

import (
	"errors"
	"os"
	"path/filepath"
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
		path, err := filepath.Abs(args[1])
		if err != nil {
			panic(err)
		}
		os.MkdirAll(path, 0755)
		return path
	}
	panic(errors.New("too many arguments"))
}
