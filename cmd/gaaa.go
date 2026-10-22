package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gaaaVerbose bool

var gaaaCmd = &cobra.Command{
	Use:     "gaaa [-v] [attr1 [attr2...]]",
	Aliases: []string{"GetAllAdminAccounts"},
	Short:   "Get All Admin Accounts",
	Long: `List the global, delegated and domain admin accounts.
With -v, display all their attributes. If attributes are given, display only those.`,

	Args: cobra.ArbitraryArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAllAdminAccounts(conf, gaaaVerbose, args)
	},
}

func init() {
	rootCmd.AddCommand(gaaaCmd)

	gaaaCmd.Flags().BoolVarP(&gaaaVerbose, "verbose", "v", false, "Display the attributes of each account")
}
