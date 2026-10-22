package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gquCmd = &cobra.Command{
	Use:     "gqu {server|id}",
	Aliases: []string{"GetQuotaUsage"},
	Short:   "Get Quota Usage",
	Long: `Display the quota usage of the accounts hosted on a mailbox server, given by name or zimbraId.
Each line shows the account name, its quota limit and its used size, in bytes. A limit of 0 means no quota.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetQuotaUsage(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(gquCmd)
}
