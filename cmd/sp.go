package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var spCmd = &cobra.Command{
	Use:     "sp {name@domain|id} password",
	Aliases: []string{"SetPassword"},
	Short:   "Set Password",
	Long: `Set the password of an account, given by name or zimbraId.
The server enforces the password policy of the account (zimbraPasswordMinLength...).`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.SetPassword(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(spCmd)
}
