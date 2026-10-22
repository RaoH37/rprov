package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var adlmCmd = &cobra.Command{
	Use:     "adlm {list@domain|id} member@domain [member@domain...]",
	Aliases: []string{"AddDistributionListMember"},
	Short:   "Add Distribution List Members",
	Long:    `Add one or more members to a distribution list, given by name or zimbraId. Members can be internal or external addresses.`,

	Args: cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.AddDistributionListMember(conf, args[0], args[1:])
	},
}

func init() {
	rootCmd.AddCommand(adlmCmd)
}
