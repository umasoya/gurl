package go_curl_wrapper

import (
	"fmt"
	"io"
	"os"
)

const version = "0.1"

const (
	ExitCodeOK = iota
	ExitCodeError
)

type GoCurlWrapper struct {
	Out,Err io.Writer
}

func (g GoCurlWrapper) Run(args []string) int {
}
