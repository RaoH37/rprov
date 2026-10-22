package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cddlCmd = &cobra.Command{
	Use:     "cddl list@domain [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateDynamicDistributionList"},
	Short:   "Create Dynamic Distribution List",
	Long: `Create a dynamic distribution list and print its zimbraId. Repeat an attribute to set several values.
The members of a dynamic list are the result of the LDAP query set in its memberURL attribute.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		args = args[1:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateDistributionList(conf, name, true, args)
	},
}

func init() {
	rootCmd.AddCommand(cddlCmd)
}
