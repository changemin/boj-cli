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
	if len(args) == 0 {
		os.Exit(1)
	}
	num, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("문제 번호를 입력해주세요")
	}

	prob := Problem{num: num}

	url := "https://www.acmicpc.net/problem/" + args[0]
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)

	prob.title = doc.Find("#problem_title").Text()
	fmt.Println("title : " + prob.title)

	prob.description = strings.TrimLeft(doc.Find("#problem_description").Text(), " ")
	fmt.Println("description : " + prob.description)

	prob.input = doc.Find("#sample-input-1").Text()
	fmt.Println("input : " + prob.input)

	prob.output = doc.Find("#sample-output-1").Text()
	fmt.Println("output : " + prob.output)

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
