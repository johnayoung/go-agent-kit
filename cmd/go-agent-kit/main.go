package main

import (
	"os"

	"github.com/johnayoung/go-agent-kit/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
