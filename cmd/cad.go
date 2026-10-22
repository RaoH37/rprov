package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cadCmd = &cobra.Command{
	Use:     "cad alias-domain {target-domain|id} [attr1 value1 [attr2 value2...]]",
	Aliases: []string{"CreateAliasDomain"},
	Short:   "Create Alias Domain",
	Long: `Create an alias domain of a target domain, given by name or zimbraId, and print its zimbraId.
Mail sent to any address of the alias domain is forwarded to the same address in the target domain.
zimbraDomainType, zimbraDomainAliasTargetId, zimbraMailCatchAllAddress and zimbraMailCatchAllForwardingAddress are set by the command.`,

	Args: cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		aliasName := args[0]
		target := args[1]

		args = args[2:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CreateAliasDomain(conf, aliasName, target, args)
	},
}

func init() {
	rootCmd.AddCommand(cadCmd)
}
