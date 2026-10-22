package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cpeCmd = &cobra.Command{
	Use:     "cpe {name@domain|id} {newName@domain}",
	Aliases: []string{"ChangePrimaryEmail"},
	Short:   "Change Primary Email",
	Long: `Rename an account, given by name or zimbraId, and keep its old name as an alias.
If the new name is already an alias of the account, that alias is removed first.
The steps are separate requests: if one fails, the error says which changes were made.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.ChangePrimaryEmail(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(cpeCmd)
}
