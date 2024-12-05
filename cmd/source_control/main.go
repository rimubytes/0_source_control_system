package main 

import (
	"fmt"
	"os"
	"github.com/rimubytes/0_source_control_system/internal/repository"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
}