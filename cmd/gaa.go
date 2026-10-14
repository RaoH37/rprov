package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gaaCmd = &cobra.Command{
	Use:     "gaa [domain]", // [domain] in brackets means optional
	Aliases: []string{"GetAllAccounts"},
	Short:   "Get All Accounts",
	Long:    `List all accounts. If a domain is given, list only the accounts of that domain.`,

	// Positional argument validation
	// At most 1 argument (the domain)
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		domain := ""
		if len(args) > 0 {
			domain = args[0]
		}

		if len(attributes) == 0 {
			attributes = []string{"mail"}
		}

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetAllAccounts(conf, domain, attributes)
	},
}

func init() {
	// Register the command on the root
	rootCmd.AddCommand(gaaCmd)

	// Flag definitions for 'gaa'
	// StringSliceVarP allows passing the same flag several times: --attr sn --attr cn
	// Or comma-separated: --attr sn,cn
	// "attributes": long name
	// "a": shorthand -a
	// []string{}: default value (empty)
	// "description"
	//gaaCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
