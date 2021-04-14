package cmd

import (
	"bj/model"
	"bj/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "get",
	Short: "문제를 파싱하여 저장합니다.",
	Long:  `여러문제, 범위 또는 다중 범위를 사용하실 수 있습니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateProblem(args)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func generateProblem(args []string) {
	if utils.ValidateConfigFile() {
		if len(args) == 0 { // 문제 번호 입력을 안했을 경우
			color.Error.Prompt("문제 번호를 입력해주세요")
			color.Green.Print("\nbj get [문제번호]")
			os.Exit(1)
		} else {
			for _, arg := range args {
				if isRange(arg) {
					offset := strings.Split(args[0], "~")
					if len(offset) > 2 {
						color.Error.Prompt("정확한 범위를 입력하세요")
						color.Green.Print("\nbj get [문제번호]~[문제번호]")
						os.Exit(1)
					}
					startNum, _ := strconv.Atoi(offset[0])
					endNum, _ := strconv.Atoi(offset[1])
					if startNum > endNum {
						color.Error.Prompt("범위는 1보다 커야 합니다.")
						color.Green.Print("\nbj get [문제번호]~[문제번호]")
						os.Exit(1)
					}
					for i := startNum; i <= endNum; i++ {
						parseProblem(i)
					}
				} else {
					num, err := strconv.Atoi(arg)
					if err != nil {
						color.Error.Prompt("문제 번호를 정수로 입력해주세요")
						color.Green.Print("\nbj get [문제번호]")
						os.Exit(1)
					}
					parseProblem(num)
				}

			}
		}

		// TODO: - table 파싱
	} else {
		color.Error.Println("설정 파일이 존재하지 않거나 완벽히 구성되지 못했습니다.")
		color.Info.Println("\nbj init 명령어로 파일을 생성하세요")
	}

}

func isRange(arg string) bool {
	return strings.Contains(arg, "~")
}

func parseProblem(num int) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	prob := model.Problem{Num: num}

	response, err := http.Get("https://www.acmicpc.net/problem/" + strconv.Itoa(num))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		color.Error.Prompt("다음 문제는 존재하지 않습니다(" + strconv.Itoa(prob.Num) + ")")
	} else {
		doc, _ := goquery.NewDocumentFromReader(response.Body)
		prob.Title = doc.Find("#problem_title").Text()
		prob.Title = strings.Replace(prob.Title, "/", "", -1) // remove `/`
		prob.Description = strings.TrimSpace(doc.Find("#problem_description").Text())
		prob.Input = strings.TrimSpace(doc.Find("#sample-input-1").Text())
		prob.Output = strings.TrimSpace(doc.Find("#sample-output-1").Text())

		if utils.IsProbExist(prob.Num) {
			color.Error.Prompt("다음 문제는 이미 존재합니다(" + strconv.Itoa(prob.Num) + ")")
		} else {
			if _, err := os.Stat(utils.GetRangeOfProb(prob.Num)); os.IsNotExist(err) {
				os.Mkdir(utils.GetRangeOfProb(prob.Num), os.ModePerm)
			}

			path := utils.GetRangeOfProb(prob.Num) + "/" + strconv.Itoa(prob.Num) + "번 - " + prob.Title

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, os.ModePerm)
			}

			f1, err := os.Create(path + "/solve" + utils.ReadFileExtension())
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
			defer f1.Close()
			color.Info.Prompt("🎉 파일 생성 성공 - " + path + "/solve" + utils.ReadFileExtension())

			fmt.Fprintf(f1, generateStrProbDescription(prob))
		}
		s.Stop()

	}
}

func generateStrProbDescription(prob model.Problem) string {
	str := ""
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, "Created By "+utils.ReadUsername()+" on "+utils.GetCurrentDate())
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, strconv.Itoa(prob.Num)+"번 : "+prob.Title)
	addStrCommentedLine(&str, "https://www.acmicpc.net/problem/"+strconv.Itoa(prob.Num))
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, "* 문제")
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, replaceNewlineWithComment(prob.Description))
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, "* 입력")
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, replaceNewlineWithComment(prob.Input))
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, "* 출력")
	addStrEmptyLine(&str)
	addStrCommentedLine(&str, replaceNewlineWithComment(prob.Output))
	addStrEmptyLine(&str)
	return str
}

func replaceNewlineWithComment(str string) string {
	str = strings.Replace(str, "\n", "\n"+utils.ReadCommentStyle()+" ", -1)
	return str
}

func addStrCommentedLine(str *string, substr string) {
	*str += utils.ReadCommentStyle() + " " + substr + "\n"
}

func addStrEmptyLine(str *string) {
	*str += utils.ReadCommentStyle() + "\n"
}
