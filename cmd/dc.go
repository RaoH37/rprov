package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var dcCmd = &cobra.Command{
	Use:     "dc {name|id}",
	Aliases: []string{"DeleteCos"},
	Short:   "Delete Cos",
	Long: `Delete a class of service (COS), given by name or zimbraId.
There is no confirmation prompt and the deletion cannot be undone.`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.DeleteCos(conf, args[0])
	},
}

func init() {
	rootCmd.AddCommand(dcCmd)
}
