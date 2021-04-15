package cmd

import (
	"bj/utils"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "맞은 문제로 표시합니다.",
	Long:  `맞은 문제로 표시합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.ReadUseReadme() {
			markProbAsSolved(args)
		} else {
			color.Error.Println("use-readme를 활성화하셔야 사용하실 수 있습니다.")
		}

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
			executeGitPush(num)
		}
	}
}

func executeGitPush(num int) {
	gitAdd := exec.Command("git", "add", ".")
	gitCommit := exec.Command("git", "commit", "-m", "Solve "+strconv.Itoa(num))
	gitPush := exec.Command("git", "push")

	output, err := gitAdd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	output, err = gitCommit.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	output, err = gitPush.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
