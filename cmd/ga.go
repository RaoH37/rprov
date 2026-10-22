package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gaCmd = &cobra.Command{
	Use:     "ga {name@domain|id} [attr1 [attr2...]]",
	Aliases: []string{"GetAccount"},
	Short:   "Get Account Attributes",
	Long: `Display the attributes of an account, given by name or zimbraId, sorted by name. If attributes are given, display only those.
The zimbraCOSId attribute is followed by the matching zimbraCOSName.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAccount(conf, key, args)
	},
}

func init() {
	rootCmd.AddCommand(gaCmd)
}
