package cmd

import (
	"fmt"
	"os"
    "agenda/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var RootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "A meeting management system base on CLI",
	Long: `A meeting management system base on CLI`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.agenda.yaml)")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" { 
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".agenda") 
	viper.AddConfigPath("$HOME") 
	viper.AutomaticEnv()         
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	service.Init()
}