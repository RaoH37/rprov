package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var rmcCmd = &cobra.Command{
	Use:     "rmc {name@domain|id}",
	Aliases: []string{"RecalculateMailboxCounts"},
	Short:   "Recalculate Mailbox Counts",
	Long: `Recalculate the quota usage and the folder counts of the mailbox of an account, given by name or zimbraId.
Print the recalculated quota usage, in bytes.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RecalculateMailboxCounts(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmcCmd)
}
