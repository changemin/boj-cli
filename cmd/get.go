package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
}

func generateProblem(num int) {
	prob := Problem{num: num}

	response, err := http.Get("https://www.acmicpc.net/problem/" + strconv.Itoa(num))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		color.Error.Prompt("❗다음 문제는 존재하지 않습니다(" + strconv.Itoa(prob.num) + ")")
	} else {
		doc, _ := goquery.NewDocumentFromReader(response.Body)
		prob.title = doc.Find("#problem_title").Text()
		prob.description = strings.TrimSpace(doc.Find("#problem_description").Text())
		prob.input = strings.TrimSpace(doc.Find("#sample-input-1").Text())
		prob.output = strings.TrimSpace(doc.Find("#sample-output-1").Text())

		makeProbDirAndFile(prob)
	}
}

func makeProbDirAndFile(prob Problem) {
	path := strconv.Itoa(prob.num) + "-" + prob.title

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	f1, err := os.Create(path + "/" + strconv.Itoa(prob.num) + ".c")
	if err != nil {
		os.Exit(1)
	}
	defer f1.Close()
	color.Info.Prompt("🎉 파일 생성 성공 - " + path + "/" + strconv.Itoa(prob.num) + ".c")

	fmt.Fprintf(f1, getProbCommentString(prob))
	fmt.Fprintf(f1, getLanguageDefaultPrintHello())
}

func getProbCommentString(prob Problem) string {
	str := ""
	str = str + "/*\n"
	str = str + getCurrentDate() + "\n\n"
	str = str + "Created By {username}\n\n"
	str = str + strconv.Itoa(prob.num) + "번 : " + prob.title + "\n"
	str = str + "https://www.acmicpc.net/problem/" + strconv.Itoa(prob.num) + "\n\n"
	str = str + "* 문제\n\n"
	str = str + prob.description + "\n\n"
	str = str + "* 입력\n\n"
	str = str + prob.input + "\n\n"
	str = str + "* 출력\n\n"
	str = str + prob.output + "\n\n"
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

func getCurrentDate() string {
	dateTime := time.Now()
	return dateTime.Format("2006-01-02")
}

// Problem 모델
type Problem struct {
	num         int
	title       string
	description string
	input       string
	output      string
	// timeLimit   string
	// memoryLimit string
	// passRatio   string
}
