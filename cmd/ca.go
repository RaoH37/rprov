package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var caCmd = &cobra.Command{
	Use:     "ca name@domain password [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateAccount"},
	Short:   "Create Account",
	Long:    `Create an account and print its zimbraId. Repeat an attribute to set several values.`,

	Args: cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		password := args[1]

		args = args[2:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateAccount(conf, name, password, args)
	},
}

func init() {
	rootCmd.AddCommand(caCmd)
}
