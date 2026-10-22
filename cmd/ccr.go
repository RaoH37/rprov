package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var ccrCmd = &cobra.Command{
	Use:     "ccr name@domain password displayName {name} zimbraCalResType {Location|Equipment} [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateCalendarResource"},
	Short:   "Create Calendar Resource",
	Long: `Create a calendar resource and print its zimbraId. Repeat an attribute to set several values.
displayName and zimbraCalResType (Location or Equipment) are required. Give "" as password to create the resource without one.`,

	Args: cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		password := args[1]

		args = args[2:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateCalendarResource(conf, name, password, args)
	},
}

func init() {
	rootCmd.AddCommand(ccrCmd)
}
