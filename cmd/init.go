package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "백준 설정파일을 생성합니다",
	Long:  `그렇대요.. (임시)`,
	Run: func(cmd *cobra.Command, args []string) {
		generateConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func generateConfigFile() {
	username := inputUsername()
	language := inputLanguage()
	f, err := os.Create(".BaekJoon.yml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "username: "+username+"\nlanguage: "+language)
	color.Info.Println("설정 파일이 생성되었습니다.")
}

func inputUsername() string {
	reader := bufio.NewReader(os.Stdin)
	color.Green.Print("이름을 입력하세요 : ")
	username, _ := reader.ReadString('\n')
	return username
}

func inputLanguage() string {
	reader := bufio.NewReader(os.Stdin)
	color.Green.Print("언어를 선택하세요\n")
	fmt.Println(`  1. C
  2. C++
  3. Java
  4. Swift
  5. Go
  6. etc`)
	color.Green.Print("(1~6) : ")
	input, _ := reader.ReadString('\n')
	languageSelection, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Println(err.Error())
		color.Error.Println("1~6 사이의 숫자를 입력해주세요")
		os.Exit(1)
	}
	switch languageSelection {
	case 1:
		return "C"
	case 2:
		return "C++"
	case 3:
		return "Java"
	case 4:
		return "Swift"
	case 5:
		return "GO"
	case 6:
		// TODO: - asking for custom language
		return "C"
	default:
		return "C"
	}

}
