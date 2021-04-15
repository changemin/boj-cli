package utils

import (
	"bj/model"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateReadme() {
	readme, err := os.Create("README.md")
	if err != nil {
		fmt.Print(err)
	}
	defer readme.Close()
	fmt.Fprintf(readme, "### 맞은 문제\n")
	fmt.Fprintf(readme, "<!--Solved-->\n")
	fmt.Fprintf(readme, "### 시도했지만 맞지 못한 문제\n")
	fmt.Fprintf(readme, "<!--Tried-->\n")
}

func AddSolvedProb(num int) {
	DeleteAllProbMarkdownLink(num)
	lines, err := Readme2Lines()
	if err != nil {
		log.Print(err.Error())
	}

	fileContent := ""
	for _, line := range lines {
		if line == "<!--Solved-->" {
			fileContent += GetProbMarkdownLink(num)
		}
		fileContent += strings.Replace(line, GetProbMarkdownLink(num), "", -1)
		fileContent += "\n"
	}
	ioutil.WriteFile("README.md", []byte(fileContent), 0644)

}

func DeleteAllProbMarkdownLink(num int) {
	lines, err := Readme2Lines()
	if err != nil {
		log.Print(err.Error())
	}

	fileContent := ""
	for _, line := range lines {
		fileContent += strings.Replace(line, strings.TrimSuffix(GetProbMarkdownLink(num), "\n"), "", -1)
		if !strings.Contains(line, strings.TrimSuffix(GetProbMarkdownLink(num), "\n")) {
			fileContent += "\n"
		}
	}
	ioutil.WriteFile("README.md", []byte(fileContent), 0644)
}

func IsProbExistInTriedProb(num int) bool {
	lines, err := Readme2Lines()
	if err != nil {
		log.Print(err.Error())
	}

	for _, line := range lines {
		if strings.Contains(line, GetProbMarkdownLink(num)) {
			return true
		}
	}
	return false
}

func GetProbMarkdownLink(num int) string {
	prob := model.Num2Prob(num)
	return "[" + strconv.Itoa(prob.Num) + "](./" + GetRangeOfProb(prob.Num) + "/" + strconv.Itoa(prob.Num) + "번%20-%20" + strings.Replace(prob.Title, " ", "%20", -1) + ") \n"
}

// [1000](./1000번~1099번/1000번%20-%20A+B)

func AddTriedProb(num int) {
	lines, err := Readme2Lines()
	if err != nil {
		log.Print(err.Error())
	}

	fileContent := ""
	for _, line := range lines {
		if line == "<!--Tried-->" {
			fileContent += GetProbMarkdownLink(num)
		}
		fileContent += line
		fileContent += "\n"
	}
	ioutil.WriteFile("README.md", []byte(fileContent), 0644)
}

func Readme2Lines() ([]string, error) {
	f, err := os.Open("./README.md")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
