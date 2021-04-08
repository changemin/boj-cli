package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "백준 설정파일을 생성합니다",
	Long:  `그렇대요.. (임시)`,
	Run: func(cmd *cobra.Command, args []string) {
		read()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func test() {
	viper.New()
	viper.SetConfigFile("bjConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("username", "NAME")
	viper.SetDefault("extension", ".c")
	viper.WriteConfig()
}

func read() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fileExtension := viper.Get("extension")
	fmt.Println(fileExtension)
	username := viper.Get("username")
	fmt.Println(username)
	placeholder := viper.GetString("placeholder")
	fmt.Println(placeholder)
}

func generateConfigFile() {
	username := inputUsername()
	fileExtension := inputFileExtension()
	f, err := os.Create(".BaekJoon.yml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "username: "+username+"\nextension: "+fileExtension)
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
