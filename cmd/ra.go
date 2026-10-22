package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var raCmd = &cobra.Command{
	Use:     "ra {name@domain|id} {newName@domain}",
	Aliases: []string{"RenameAccount"},
	Short:   "Rename Account",
	Long: `Rename an account, given by name or zimbraId. The zimbraId and the mailbox are kept.
The domain of the new name must exist.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RenameAccount(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(raCmd)
}
