package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func testProblem(problemNumber uint8) {
	fmt.Printf("test Problem: %c\n", problemNumber)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your code",
	Long:  "Pass test case to your source code.",
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
			testProblem(problemNumber)
		}
	},
}

func init() {
	testCmd.PersistentFlags().BoolP("help", "h", false, "help for enter")
	_ = testCmd.PersistentFlags().MarkHidden("help")

	testCmd.SetUsageTemplate(
		strings.Replace(
			testCmd.UsageTemplate(),
			"{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}",
			"{{.UseLine}} [problem number]{{end}}{{if .HasAvailableSubCommands}}",
			1,
		),
	)
}
