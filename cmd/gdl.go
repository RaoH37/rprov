package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gdlCmd = &cobra.Command{
	Use:     "gdl name [attr1 [attr2...]]",
	Aliases: []string{"GetDistributionList"},
	Short:   "Get DistributionList Attributes",
	Long:    `Display the attributes of a distribution list, sorted by name. If attributes are given, display only those.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetDistributionList(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gdlCmd)
}
