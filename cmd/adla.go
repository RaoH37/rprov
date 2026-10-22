package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var adlaCmd = &cobra.Command{
	Use:     "adla {list@domain|id} {alias@domain}",
	Aliases: []string{"AddDistributionListAlias"},
	Short:   "Add Distribution List Alias",
	Long:    `Add an alias to a distribution list, given by name or zimbraId. The alias domain must exist.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.AddDistributionListAlias(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(adlaCmd)
}
