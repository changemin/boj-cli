package cmd

import (
	"bj/utils"
	"os"
	"strconv"

	"github.com/gookit/color"
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
	if len(args) == 0 { // 문제 번호 입력을 안했을 경우
		color.Error.Prompt("문제 번호를 입력해주세요")
		color.Green.Print("\nbj solve [문제번호]")
		os.Exit(1)
	} else {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			color.Error.Prompt("문제 번호를 정수로 입력해주세요")
			color.Green.Print("\nbj solve [문제번호]")
			os.Exit(1)
		} else {
			utils.AddSolvedProb(num)
			// commit
		}
	}
}
