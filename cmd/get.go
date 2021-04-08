package cmd

import (
	model "bj/model"
	utils "bj/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "get",
	Short: "백준 문제를 파싱하여 저장합니다.",
	Long: `1. bj get [문제번호] : 문제번호의 문제를 가져옵니다
2. bj get [문제번호] [문제번호] [문제번호] : 여러문제를 한번에 가져옵니다
3. bj get [문제번호]~[문제번호] : 범위 내의 문제를 가져옵니다`,
	Run: func(cmd *cobra.Command, args []string) {
		parseProblem(args)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func parseProblem(args []string) {
	if utils.IsConfigFileExist() {
		if len(args) == 0 { // 문제 번호 입력을 안했을 경우
			color.Error.Prompt("문제 번호를 입력해주세요")
			color.Green.Print("\nbj get [문제번호]")
			os.Exit(1)
		} else if strings.Contains(args[0], "~") {
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
				generateProblem(i)
			}
		} else {
			for _, strProbNum := range args {
				num, err := strconv.Atoi(strProbNum)
				if err != nil {
					color.Error.Prompt("문제 번호를 정수로 입력해주세요")
					color.Green.Print("\nbj get [문제번호]")
					os.Exit(1)
				}
				generateProblem(num)
			}
		}

		// TODO: - table 파싱
	} else {
		color.Error.Println("설정 파일이 존재하지 않습니다.")
		color.Info.Println("\nbj init 명령어로 파일을 생성하세요")
	}

}

func generateProblem(num int) {
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
		prob.Description = strings.TrimSpace(doc.Find("#problem_description").Text())
		prob.Input = strings.TrimSpace(doc.Find("#sample-input-1").Text())
		prob.Output = strings.TrimSpace(doc.Find("#sample-output-1").Text())

		makeProbDirAndFile(prob)
	}
}

func makeProbDirAndFile(prob model.Problem) {
	if isProbExist(prob) {
		color.Error.Prompt("다음 문제는 이미 존재합니다(" + strconv.Itoa(prob.Num) + ")")
	} else {
		if _, err := os.Stat(getStrRangeOfProb(prob.Num)); os.IsNotExist(err) {
			os.Mkdir(getStrRangeOfProb(prob.Num), os.ModePerm)
		}

		path := getStrRangeOfProb(prob.Num) + "/" + strconv.Itoa(prob.Num) + "번 - " + prob.Title

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}

		f1, err := os.Create(path + "/solve.c")
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		defer f1.Close()
		color.Info.Prompt("🎉 파일 생성 성공 - " + path + "/solve.c")

		fmt.Fprintf(f1, getProbCommentString(prob))
		fmt.Fprintf(f1, getLanguageDefaultPrintHello())
	}
}

func isProbExist(prob model.Problem) bool {
	rangeFolderList, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, rangeFolder := range rangeFolderList {
		if rangeFolder.Name() == getStrRangeOfProb(prob.Num) {
			files, err := ioutil.ReadDir(getStrRangeOfProb(prob.Num))
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				if strings.Contains(file.Name(), strconv.Itoa(prob.Num)) {
					if filerc, _ := os.Open(getStrRangeOfProb(prob.Num) + "/" + file.Name() + "/" + strconv.Itoa(prob.Num) + ".c"); filerc != nil {
						return true
					}
				}
			}

		}

	}

	return false
}

func getProbCommentString(prob model.Problem) string {
	str := ""
	str = str + "/*\n"
	str = str + utils.GetCurrentDate() + "\n\n"
	str = str + "Created By " + utils.ReadUsername() + "\n\n"
	str = str + strconv.Itoa(prob.Num) + "번 : " + prob.Title + "\n"
	str = str + "https://www.acmicpc.net/problem/" + strconv.Itoa(prob.Num) + "\n\n"
	str = str + "* 문제\n\n"
	str = str + prob.Description + "\n\n"
	str = str + "* 입력\n\n"
	str = str + prob.Input + "\n\n"
	str = str + "* 출력\n\n"
	str = str + prob.Output + "\n\n"
	str = str + "*/\n\n"
	return str
}

func getLanguageDefaultPrintHello() string {
	return `#include<stdio.h>

	int main() {
		printf("Hello, World!");

		return 0;
	}`

}

func getStrRangeOfProb(num int) string {
	strNum := strconv.Itoa(num)
	return strNum[:len(strNum)-2] + "00번~" + strNum[:len(strNum)-2] + "99번"
}
