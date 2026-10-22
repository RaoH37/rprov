package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var rdlmCmd = &cobra.Command{
	Use:     "rdlm {list@domain|id} member@domain [member@domain...]",
	Aliases: []string{"RemoveDistributionListMember"},
	Short:   "Remove Distribution List Members",
	Long:    `Remove one or more members from a distribution list, given by name or zimbraId.`,

	Args: cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RemoveDistributionListMember(conf, args[0], args[1:])
	},
}

func init() {
	rootCmd.AddCommand(rdlmCmd)
}
