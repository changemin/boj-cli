package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bj",
	Short: "BaekJoon-CLI",
	Long: `백준 문제풀이 및 파일 관리를 도와줍니다
	
			https://github.com/Changemin/BaekJoon-CLI`,

	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("hello")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bj.yaml)")

	viper.SetDefault("username", "USERNAME")
	viper.SetDefault("language", "C")

	// viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bj" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bj")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
