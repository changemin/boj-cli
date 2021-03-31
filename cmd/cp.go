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
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "작성한 코드를 클립보드에 카피합니다.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		copyCode2Clipboard(args)
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
}

func copyCode2Clipboard(args []string) {
	fmt.Printf(Green, "📋 클립보드에 복사되었습니다!")
	clipboard.WriteAll("hello")
}
