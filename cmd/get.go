package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "get",
	Short: "문제를 다운 파싱합니다",
	Long:  `문제 번호를 입력(미작성)`,
	Run: func(cmd *cobra.Command, args []string) {
		parseProblem(args)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func parseProblem(args []string) {
	if len(args) == 0 { // 문제 번호 입력을 안했을 경우
		fmt.Printf(Green, "문제 번호를 입력해주세요\n\nbj get [문제번호]")
		os.Exit(1)
	}
	for idx, strProbNum := range args {
		num, err := strconv.Atoi(strProbNum)

		if err != nil {
			fmt.Printf(Green, "문제 번호 정수로 입력해주세요\n\nbj get [문제번호]")
			os.Exit(1)
		}

		prob := Problem{num: num}

		url := "https://www.acmicpc.net/problem/" + strProbNum
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)

		fmt.Println(strconv.Itoa(idx) + "Prob num : " + strconv.Itoa(prob.num))
		prob.title = doc.Find("#problem_title").Text()
		fmt.Println("title : " + prob.title)

		prob.description = strings.TrimSpace(doc.Find("#problem_description").Text())
		fmt.Println("description : " + prob.description)

		prob.input = strings.TrimSpace(doc.Find("#sample-input-1").Text())
		fmt.Println("input : " + prob.input)

		prob.output = strings.TrimSpace(doc.Find("#sample-output-1").Text())
		fmt.Println("output : " + prob.output)

		makeProbDirAndFile(prob)
	}

	// TODO: - table 파싱
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
	fmt.Fprintf(f1, getProbCommentString(prob))
	fmt.Fprintf(f1, getLanguageDefaultPrintHello())
}

func getProbCommentString(prob Problem) string {
	str := ""
	str = str + "/*\n"
	str = str + "2020-03-31\n\n"
	str = str + "Created By {username}\n\n"
	str = str + strconv.Itoa(prob.num) + "번 : " + prob.title + "\n"
	str = str + "https://www.acmicpc.net/problem/" + strconv.Itoa(prob.num) + "\n\n"
	str = str + "* 문제\n\n"
	str = str + "\t" + prob.description + "\n\n"
	str = str + "* 입력\n\n"
	str = str + "\t" + prob.input + "\n\n"
	str = str + "* 출력\n\n"
	str = str + "\t" + prob.output + "\n\n"
	str = str + "*/\n\n"
	return str
}

func getLanguageDefaultPrintHello() string {
	return `#include<stdio.h>

int main() {
	printf("Hello, World!")
}`
}

func appendLineWithNewLine(str string, to string) string {
	return to + str + "\n"
}

// func makeProbFile

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

const (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)
