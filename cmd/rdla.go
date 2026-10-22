package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var rdlaCmd = &cobra.Command{
	Use:     "rdla {list@domain|id} {alias@domain}",
	Aliases: []string{"RemoveDistributionListAlias"},
	Short:   "Remove Distribution List Alias",
	Long:    `Remove an alias from a distribution list, given by name or zimbraId.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RemoveDistributionListAlias(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(rdlaCmd)
}
