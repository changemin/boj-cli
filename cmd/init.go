package cmd

import (
	utils "bj/utils"
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
	useReadme := inputUseReadme()
	f, err := os.Create("bjConfig.yaml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "username: "+username)
	fmt.Fprintf(f, "file-extension: "+fileExtension)
	fmt.Fprintf(f, "comment-style: \""+strings.TrimSpace(commentStyle)+"\"\n")
	if useReadme == true {
		fmt.Fprintf(f, "use-readme: true")
		if utils.ReadUseReadme() {
			utils.CreateReadme()
		}
	} else {
		fmt.Fprintf(f, "use-readme: false")
	}

	color.Info.Println("\n🎉 설정 파일이 생성되었습니다.")
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
		color.Green.Println("\n파일 확장자를 입력해주세요 ex) .c, .java")
		color.Green.Print(">>> ")
		input, _ := reader.ReadString('\n')
		if strings.Contains(input, ".") {
			return input
		} else {
			color.Info.Println("\n.c, .java 와 같은 형식이어야 합니다.")
		}
	}
	return ".c"
}

func inputUseReadme() bool {
	reader := bufio.NewReader(os.Stdin)
	for true {
		color.Green.Println("\nReadme를 사용하시겠습니까? (y/n)")
		color.Green.Print(">>> ")
		input, _ := reader.ReadString('\n')
		if input == "y\n" || input == "Y\n" {
			return true
		} else if input == "n\n" || input == "N\n" {
			return false
		} else {
			color.Info.Println("y 또는 n을 입력해주세요")
		}
	}
	return false
}

func inputCommentStyle() string {
	reader := bufio.NewReader(os.Stdin)
	color.Green.Println("\n주석 형식을 입력해주세요 ex) //, #")
	color.Green.Print(">>> ")
	commentStyle, _ := reader.ReadString('\n')
	return commentStyle
}
