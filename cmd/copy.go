package cmd

import (
	utils "bj/utils"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "작성한 코드를 클립보드에 카피합니다.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		copyCode2Clipboard(args)
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
}

func copyCode2Clipboard(args []string) {
	if len(args) == 0 { // 문제 번호 입력을 안했을 경우
		color.Error.Prompt("문제 번호를 입력해주세요")
		color.Green.Print("\nbj cp [문제번호]")
		os.Exit(1)
	} else if len(args) > 1 {
		color.Error.Prompt("한개의 문제 번호만 입력해주세요")
		color.Green.Print("\nbj cp [문제번호]")
		os.Exit(1)
	} else {
		rangeFolderList, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		probNum, _ := strconv.Atoi(args[0])

		for _, rangeFolder := range rangeFolderList {
			if rangeFolder.Name() == utils.GetRangeOfProb(probNum) {
				files, err := ioutil.ReadDir(utils.GetRangeOfProb(probNum))
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range files {
					if strings.Contains(file.Name(), strconv.Itoa(probNum)) {
						filerc, err := os.Open(utils.GetRangeOfProb(probNum) + "/" + file.Name() + "/solve" + utils.ReadFileExtension())
						if err != nil {
							log.Fatal(err)
						}
						defer filerc.Close()
						buf := new(bytes.Buffer)
						buf.ReadFrom(filerc)
						contents := buf.String()
						clipboard.WriteAll(contents)
						color.Info.Println("📋 '" + file.Name() + "'이(가) 클립보드에 복사되었습니다!")
						os.Exit(1)
					}
				}

			}

		}

		color.Error.Prompt("❗다음 문제는 존재하지 않습니다(" + args[0] + ")")
		os.Exit(1)
	}
}
