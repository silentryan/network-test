package input

import (
	"fmt"
	// "os"
)

// Handle Args: Need sort or combine args to a clearly command
// Examples:
// `-h host` and `-f path` should not come together
// -A: Accessible testing
// -L: Latency testing

// Open An CMD and Receive User's command
// According user's command and excute different function
func Cmd() {}

// execute single command
func Excute(arg string) error {
	// 1. Check if arg on support function list
	// if not, return error: function temporarily not support and standard support statment

	// 2. According to arg, run functions
	switch arg {
	case "-f":
		fmt.Println("Will handling your files")
	case "-h":
		fmt.Println("Testing your host")
	}

	return nil
}
