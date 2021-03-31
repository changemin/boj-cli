package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "백준 WorkSpace를 생성합니다",
	Long:  `그렇대요.. (임시)`,
	Run: func(cmd *cobra.Command, args []string) {
		generateConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func generateConfigFile() {
	fmt.Println("generate init file if not exist")
	f, err := os.Create(".bj.yaml")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	fmt.Fprintf(f, string("username: USERNAME\nlanguage: C"))
}
