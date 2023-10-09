package cmd

import (
	"CFC/python"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func enterContest(contestNumber int) {
	python.RunPy("enter.py", strconv.Itoa(contestNumber))
}

var enterCmd = &cobra.Command{
	Use:   "enter",
	Short: "Enter contest",
	Long:  "Enter contest. Give contest number.",
	Run: func(cmd *cobra.Command, args []string) {
		argN := len(args)
		if argN == 0 {
			fmt.Println(
				"Missing contest number.\n" +
					"Use \"cfc --help\" for more information about a command.",
			)
		} else if argN > 1 {
			fmt.Println(
				"Can't have multiple arguments after 'cfc enter'\n" +
					"Use \"cfc --help\" for more information about a command.",
			)
		} else {
			contestNumber, err := strconv.Atoi(args[0])
			if err == nil {
				enterContest(contestNumber)
			} else {
				fmt.Println(
					"Contest number should be a single integer number.\n" +
						"Use \"cfc --help\" for more information about a command.",
				)
			}
		}
	},
}

func init() {
	enterCmd.PersistentFlags().BoolP("help", "h", false, "help for enter")
	_ = enterCmd.PersistentFlags().MarkHidden("help")

	enterCmd.SetUsageTemplate(
		strings.Replace(
			enterCmd.UsageTemplate(),
			"{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}",
			"{{.UseLine}} [contest number]{{end}}{{if .HasAvailableSubCommands}}",
			1,
		),
	)
}
