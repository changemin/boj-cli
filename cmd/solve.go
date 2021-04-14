package cmd

import (
	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "맞은 문제로 표시합니다.",
	Long:  `맞은 문제로 표시합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		markProbAsSolved(args)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
}

func markProbAsSolved(args []string) {

}
