package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var dcrCmd = &cobra.Command{
	Use:     "dcr {name@domain|id}",
	Aliases: []string{"DeleteCalendarResource"},
	Short:   "Delete Calendar Resource",
	Long: `Delete a calendar resource, given by name or zimbraId, along with its mailbox.
There is no confirmation prompt and the deletion cannot be undone.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.DeleteCalendarResource(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(dcrCmd)
}
