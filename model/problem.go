package model

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
)

// Problem 모델
type Problem struct {
	Num         int
	Title       string
	Description string
	Input       string
	Output      string
	// timeLimit   string
	// memoryLimit string
	// passRatio   string
}

func Num2Prob(num int) Problem {
	prob := Problem{Num: num}

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
		return prob
	}
	return prob
}
