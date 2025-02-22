package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/reeflective/console"
)

func ExitConsole(c *console.Console) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Confirm exit (y/n): ")
	text, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(text)

	if (answer == "Y") || (answer == "y") {
		return true
	}
	return false
}
