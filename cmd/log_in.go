package cmd

import (
	"agenda/service"
	"github.com/spf13/cobra"
)

var log_inCmd = &cobra.Command{
	Use:   "log_in -n [username] -p [password]",
	Short: "log in",
	Long: `Input command mode like : log_in -n Go -p 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		service.Log_in(tmp_n, tmp_p)
	},
}

func init() {
	RootCmd.AddCommand(log_inCmd)
	log_inCmd.Flags().StringP("name", "n", "", "user name")
	log_inCmd.Flags().StringP("password", "p", "", "user password")
}