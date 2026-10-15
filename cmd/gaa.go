package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gaaCmd = &cobra.Command{
	Use:     "gaa [domain]",
	Aliases: []string{"GetAllAccounts"},
	Short:   "Get All Accounts",
	Long:    `List all accounts. If a domain is given, list only the accounts of that domain.`,

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		domain := ""
		if len(args) > 0 {
			domain = args[0]
		}

		if len(attributes) == 0 {
			attributes = []string{"mail"}
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAllAccounts(conf, domain, attributes)
	},
}

func init() {
	rootCmd.AddCommand(gaaCmd)
}
