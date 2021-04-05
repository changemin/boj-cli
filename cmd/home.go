package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:   "home",
	Short: ``,
	Run: func(cmd *cobra.Command, args []string) {
		goHome()
	},
}

func init() {
	rootCmd.AddCommand(homeCmd)
}

func goHome() {
	err := os.Chdir("./1000번~1099번")
	log.Print(err)
}
