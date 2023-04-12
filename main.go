package main

import (
	"fmt"
	"os"

	"github.com/jatalocks/terracove/cmd"
)

var version = "0.0.5"

func main() {
	if err := cmd.Execute(version, false); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
