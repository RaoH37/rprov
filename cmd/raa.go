package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var raaCmd = &cobra.Command{
	Use:     "raa {name@domain|id} {alias@domain}",
	Aliases: []string{"RemoveAccountAlias"},
	Short:   "Remove Account Alias",
	Long:    `Remove an alias from an account, given by name or zimbraId.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RemoveAccountAlias(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(raaCmd)
}
