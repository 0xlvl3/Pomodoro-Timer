package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	r := lipgloss.DefaultRenderer()
	s := r.NewStyle().Foreground(lipgloss.Color("205"))
	max := 10
	for i := max; i > 0; i-- {
		fmt.Printf(s.Render("\rStudy Countdown: %d "), i)
		time.Sleep(time.Second * 1)
	}
}
