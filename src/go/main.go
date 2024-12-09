package main

import (
	"GoPack/src/go/pkg"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var titleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true).
	Italic(true)

func printHelp() {
	title := titleStyle.Render("Go Project Demo")
	fmt.Println(title)
	fmt.Println("Ussage: main.gi")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "one":
		pkg.One()
	case "two":
		pkg.Two()
	case "--help", "help":
		printHelp()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}
}
