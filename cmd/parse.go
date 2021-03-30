/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// parseProblemCmd represents the parseProblem command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called")
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseProblemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseProblemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// printParse()
	parseProblem(1000)
}

func parseProblem(num int) {
	response, err := http.Get("https://www.acmicpc.net/problem/1000")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	prob := Problem{num: num}

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

	// doc.Find("table tr td").Each(func(i int, td *goquery.Selection) {
	// 	fmt.Println(td)
	// })

	// fmt.Println("timeLimit : " + prob.output)

	// prob.memoryLimit = doc.Find("#sample-output-1").Text()
	// fmt.Println(prob.output)

	// prob.passRatio = doc.Find("#sample-output-1").Text()
	// fmt.Println(prob.output)

}

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

func printParse() {
	resp, err := http.Get("https://www.naver.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// title 태그의 적힌 text를 출력
	title := doc.Find("title").Text()
	fmt.Println(title)

	// html tag의 id가 news_cast인 태그의 text를 출력
	news_cast := doc.Find("#news_cast").Text()
	fmt.Println(news_cast)

	// div중 3번째 있는 div가져와 text 출력
	div_three, _ := doc.Find("div").Eq(2).Html()
	fmt.Println(div_three)

	// 함수를 전달해 사용할 수도 있다. 모든 div의 번호와 내용을 출력
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strconv.Itoa(i) + "번째 div" + s.Text())
	})
}
