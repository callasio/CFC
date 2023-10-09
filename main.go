package main

import (
	"CFC/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		println("Please give valid arguments.")
		println("Use \"cfc --help\" for more information about a command.")
		return
	}
}
