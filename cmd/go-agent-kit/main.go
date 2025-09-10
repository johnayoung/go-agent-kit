package main

import (
	"github.com/johnayoung/go-agent-kit/internal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
