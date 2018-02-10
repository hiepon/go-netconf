// -*- coding: utf-8 -*-

package main

import (
	"os"

	cmd "github.com/hiepon/go-netconf/ncclient/cmd"
)

func main() {
	if err := cmd.RootCmd("ncclient").Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
