package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gdCmd = &cobra.Command{
	Use:     "gd name [attr1 [attr2...]]",
	Aliases: []string{"GetDomain"},
	Short:   "Get Domain Attributes",
	Long: `Display the attributes of a domain, sorted by name. If attributes are given, display only those.
The zimbraDomainDefaultCOSId attribute is followed by the matching zimbraDomainDefaultCOSName.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetDomain(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gdCmd)
}
