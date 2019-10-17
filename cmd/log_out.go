package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

var log_outCmd = &cobra.Command{
	Use:   "log_out",
	Short: "log out",
	Long: `Input command mode like : log_out`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.GetFlag() == true {
			service.Log_out()
		} else {
			fmt.Println("You don't log in!")
		}		
	},
}

func init() {
	RootCmd.AddCommand(log_outCmd)
}