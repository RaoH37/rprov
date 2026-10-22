package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var ddCmd = &cobra.Command{
	Use:     "dd {domain|id}",
	Aliases: []string{"DeleteDomain"},
	Short:   "Delete Domain",
	Long: `Delete a domain, given by name or zimbraId. The server refuses to delete a domain that still has accounts, lists or aliases.
There is no confirmation prompt and the deletion cannot be undone.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.DeleteDomain(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(ddCmd)
}
