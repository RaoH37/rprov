package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cdlCmd = &cobra.Command{
	Use:     "cdl list@domain [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateDistributionList"},
	Short:   "Create Distribution List",
	Long:    `Create a distribution list and print its zimbraId. Repeat an attribute to set several values.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		args = args[1:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateDistributionList(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(cdlCmd)
}
