package main

import (
	"os"

	"github.com/umasoya/gurl"
)

func main() {
	gurl := gurl.GoCurlWrapper{Out: os.Stdout, Err: os.Stderr}
	exitCode := gurl.Run(os.Args[1:])
	os.Exit(exitCode)
}
