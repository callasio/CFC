package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func submitProblem(problemNumber uint8) {
	fmt.Printf("submit Problem: %c\n", problemNumber)
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit your code",
	Long:  "Submit your source code. (GNU G++17 7.3.0)",
	Run: func(cmd *cobra.Command, args []string) {
		argN := len(args)
		if argN == 0 {
			fmt.Println(
				"Missing problem number.\n" +
					"Use \"cfc --help\" for more information about a command.",
			)
		} else if argN > 1 {
			fmt.Println(
				"Can't have multiple arguments after 'cfc enter'\n" +
					"Use \"cfc --help\" for more information about a command.",
			)
		} else if len(args[0]) != 1 {
			fmt.Println(
				"Problem number should be a single character.\n" +
					"Use \"cfc --help\" for more information about a command.",
			)
		} else {
			problemNumber := args[0][0]
			submitProblem(problemNumber)
		}
	},
}

func init() {
	submitCmd.PersistentFlags().BoolP("help", "h", false, "help for enter")
	_ = submitCmd.PersistentFlags().MarkHidden("help")

	submitCmd.SetUsageTemplate(
		strings.Replace(
			submitCmd.UsageTemplate(),
			"{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}",
			"{{.UseLine}} [problem number]{{end}}{{if .HasAvailableSubCommands}}",
			1,
		),
	)
}
