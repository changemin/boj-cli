package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bj",
	Short: "BaekJoon-CLI",
	Long: `백준 문제풀이 및 파일 관리를 도와줍니다
	
			https://github.com/Changemin/BaekJoon-CLI`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// fileExtension := viper.Get("extension")
	// fmt.Println(fileExtension)
	// username := viper.Get("username")
	// fmt.Println(username)
}
