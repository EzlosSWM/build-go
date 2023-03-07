package main

import (
	"flag"
	"os"

	"github.com/EzlosSWM/build-go/cmd"
)

// make concurrent
func main() {

	// setup subcommand
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	// flag
	newDefault := newCmd.Bool("default", true, "Create go")

	// wash args
	if len(os.Args) < 2 {
		newCmd.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		handleActions(newCmd, newDefault)
	}

}

func handleActions(c *flag.FlagSet, newDefault *bool) {
	if *newDefault {
		cmd.AllActions()
	}
}
