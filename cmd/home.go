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
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "백준 WorkSpace의 홈으로 이동합니다",
	Long:  `미작성`,
	Run: func(cmd *cobra.Command, args []string) {
		// goHome()
		test()
	},
}

func init() {
	rootCmd.AddCommand(homeCmd)
}

func goHome() {
	home, _ := os.UserHomeDir()
	err := os.Chdir(filepath.Join(home, "Repositories"))
	if err != nil {
		panic(err)
	}
}

func test() {
	cmd := exec.Command("cd ~")
	cmd.Start()
	cmd.Wait()
	// cmd = exec.Command("sh", "-c", "echo hello,world!")
	// stdoutStderr, _ := cmd.CombinedOutput()
	// fmt.Println(string(stdoutStderr))
}
