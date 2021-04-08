package cmd

import (
	"bj/utils"
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		openProbFolder(args)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func openProbFolder(args []string) {
	if len(args) == 0 { // 문제 번호 입력을 안했을 경우
		color.Error.Prompt("문제 번호를 입력해주세요")
		color.Green.Print("\nbj get [문제번호]")
		os.Exit(1)
	} else {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			color.Error.Prompt("문제 번호를 정수로 입력해주세요")
			color.Green.Print("\nbj get [문제번호]")
			os.Exit(1)
		} else {
			if utils.IsProbExist(num) {
				open.Run("./" + utils.GetRangeOfProb(num))
			} else {
				color.Error.Println("다음 문제는 존재하지 않습니다.(" + args[0] + ")")
			}

		}
	}
}
