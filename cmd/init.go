package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "BOJ CLI 설정파일을 생성합니다.",
	Long:  `bjConfig.yaml 파일을 생성합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func generateConfigFile() {
	username := inputUsername()
	fileExtension := inputFileExtension()
	commentStyle := inputCommentStyle()
	f, err := os.Create("bjConfig.yaml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "username: "+username+"file-extension: "+fileExtension+"comment-style: \""+strings.TrimSpace(commentStyle)+"\"")
	color.Info.Println("설정 파일이 생성되었습니다.")
}

func inputUsername() string {
	reader := bufio.NewReader(os.Stdin)
	color.Green.Print("이름을 입력하세요 : ")
	username, _ := reader.ReadString('\n')
	return username
}

func inputFileExtension() string {
	reader := bufio.NewReader(os.Stdin)
	for true {
		color.Green.Println("파일 확장자를 입력해주세요 ex) .c, .java")
		color.Green.Print(">>> ")
		input, _ := reader.ReadString('\n')
		if strings.Contains(input, ".") {
			return input
		} else {
			color.Info.Println(".c, .java 와 같은 형식이어야 합니다.")
		}
	}
	return ".c"
}

func inputCommentStyle() string {
	reader := bufio.NewReader(os.Stdin)
	color.Green.Println("주석 형식을 입력해주세요 ex) //, #")
	color.Green.Print(">>> ")
	commentStyle, _ := reader.ReadString('\n')
	return commentStyle
}
