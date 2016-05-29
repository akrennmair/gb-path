package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/constabulary/gb/cmd"
)

func main() {
	vendorPath := true

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "simple":
			vendorPath = false
		case "-h", "help":
			usage()
		}
	}

	projectroot := os.Getenv("GB_PROJECT_DIR")
	if projectroot == "" {
		fmt.Fprintf(os.Stderr, "FATAL: don't run this binary directly, it is meant to be run as 'gb path'\n")
		os.Exit(1)
	}

	root, err := cmd.FindProjectroot(projectroot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: could not locate project root: %v\n", err)
		os.Exit(1)
	}

	envValue := root

	if vendorPath {
		envValue += ":" + filepath.Join(root, "vendor")
	}

	fmt.Printf("GOPATH=%q; export GOPATH;\n", envValue)
}

func usage() {
	fmt.Fprintf(os.Stderr, `gb-path, a gb plugin to find out and set your $GOPATH.

Usage:

	gb path [simple]

This will print a short shell snippet that sets your $GOPATH. It can be used 
in conjunction with the "eval" shell command:

	$ eval `+"`gb path`"+`

You can also set up an alias:

	$ alias gbpath='eval `+"`gb path`"+`

If you don't want the vendor directory to be included in $GOPATH, call gb-path 
with the parameter "simple".
`)

	os.Exit(0)
}
