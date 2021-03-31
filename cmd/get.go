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
		fmt.Println("title :\n" + prob.title)

		prob.description = strings.TrimLeft(doc.Find("#problem_description").Text(), " ")
		fmt.Println("description :\n" + prob.description)

		prob.input = doc.Find("#sample-input-1").Text()
		fmt.Println("input :\n" + prob.input)

		prob.output = doc.Find("#sample-output-1").Text()
		fmt.Println("output :\n" + prob.output)

		path := strconv.Itoa(prob.num) + "-" + prob.title

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
	}

	// TODO: - table 파싱
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
