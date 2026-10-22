package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var rdlCmd = &cobra.Command{
	Use:     "rdl {list@domain|id} {newName@domain}",
	Aliases: []string{"RenameDistributionList"},
	Short:   "Rename Distribution List",
	Long: `Rename a distribution list, given by name or zimbraId. Its zimbraId and members are kept.
The domain of the new name must exist.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RenameDistributionList(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(rdlCmd)
}
