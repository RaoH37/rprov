package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var daCmd = &cobra.Command{
	Use:     "da {name@domain|id}",
	Aliases: []string{"DeleteAccount"},
	Short:   "Delete Account",
	Long: `Delete an account, given by name or zimbraId, along with its mailbox.
There is no confirmation prompt and the deletion cannot be undone.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.DeleteAccount(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(daCmd)
}
