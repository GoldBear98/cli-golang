package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [phone]",
	Short: "Register a new user",
	Long: `Input command model like: register -n Golang -p 123456 -e 12@qq.com -t 12580`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		tmp_e, _ := cmd.Flags().GetString("email")
		tmp_t, _ := cmd.Flags().GetString("phone")
		if service.GetFlag() == false {
			service.RegisterUser(tmp_n, tmp_p,tmp_e,tmp_t)
		} else {
			fmt.Println("You have already logged in. Please log out first, and then you can regist!")
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("name", "n", "", "user name")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("phone", "t", "", "user phone")
}