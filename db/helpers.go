package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// main loop handle case of looping
func TimeLoop(label string, minutes int) {
	fmt.Printf("%s starting for %d minutes .. ", label, minutes)

	for i := minutes; i >= 0; i-- {
		fmt.Printf("\r%s Break Countdown...: %d ", label, i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}
