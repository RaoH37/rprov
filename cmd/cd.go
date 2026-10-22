package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:     "cd domain [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateDomain"},
	Short:   "Create Domain",
	Long:    `Create a domain and print its zimbraId. Repeat an attribute to set several values.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		args = args[1:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateDomain(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
