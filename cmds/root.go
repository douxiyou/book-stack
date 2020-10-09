package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "show",
	Short: "介绍自己",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("默认执行")
	},
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("不可接受的错误")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(prepare)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	ex, osErr := os.Executable()
	if osErr != nil {
		fmt.Println(fmt.Errorf("os err: %s \n", osErr))
		os.Exit(1)
	}
	viper.AddConfigPath(filepath.Dir(ex))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Read config file error: %s \n", err))
		os.Exit(1)
	}
}

func prepare() {
	fmt.Println("预备.....")
}
