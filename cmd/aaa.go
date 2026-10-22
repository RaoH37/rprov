package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var aaaCmd = &cobra.Command{
	Use:     "aaa {name@domain|id} {alias@domain}",
	Aliases: []string{"AddAccountAlias"},
	Short:   "Add Account Alias",
	Long:    `Add an alias to an account, given by name or zimbraId. The alias domain must exist.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.AddAccountAlias(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(aaaCmd)
}
