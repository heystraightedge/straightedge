package main

import (
	"os"

	"github.com/heystraightedge/straightedge/cmd/strd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
