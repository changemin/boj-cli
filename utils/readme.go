package utils

import (
	"bj/model"
	"fmt"
	"os"
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

func AddSolvedProb(prob model.Problem) {

}

func AddTriedProb(prob model.Problem) {

}
