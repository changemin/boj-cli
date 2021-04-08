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
	fileExtension := inputFileExtension()
	commentStyle := inputCommentStyle()
	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "username: "+username+"file-extension: "+fileExtension+"comment-style: "+commentStyle)
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
		color.Green.Print("파일 확장자를 입력해주세요 ex) .c, .java")
		color.Green.Print("\n>>> ")
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
	color.Green.Print("주석 : ")
	commentStyle, _ := reader.ReadString('\n')
	return commentStyle
}
