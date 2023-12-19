package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/timchurchard/bitcoin-puzzles/cmd"
)

const cliName = "bitcoin-puzzles"

func main() {
	if len(os.Args) < 2 {
		usageRoot()
	}

	// Save the command and reset the flags
	command := os.Args[1]
	flag.CommandLine = flag.NewFlagSet(cliName, flag.ExitOnError)
	os.Args = append([]string{cliName}, os.Args[2:]...)

	switch command {
	case "32btc":
		os.Exit(cmd.BTC32(os.Stdout))
	case "bobby":
		os.Exit(cmd.Bobby(os.Stdout))
	}

	usageRoot()
}

func usageRoot() {
	fmt.Printf("usage: %s command(32btc|bobby) options\n", cliName)
	os.Exit(1)
}
