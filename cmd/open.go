package cmd

import (
	"bj/utils"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "문제가 저장된 폴더를 엽니다.",
	Long:  `bj open [문제번호]로 폴더를 엽니다.`,
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
				files, err := ioutil.ReadDir(utils.GetRangeOfProb(num))
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range files {
					if strings.Contains(file.Name(), strconv.Itoa(num)) {
						if filerc, _ := os.Open(utils.GetRangeOfProb(num) + "/" + file.Name() + "/solve" + utils.ReadFileExtension()); filerc != nil {
							open.Run("./" + utils.GetRangeOfProb(num) + "/" + file.Name())
						}
					}
				}

			} else {
				color.Error.Println("다음 문제는 존재하지 않습니다.(" + args[0] + ")")
			}

		}
	}
}
