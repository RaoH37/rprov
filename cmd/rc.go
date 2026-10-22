package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var rcCmd = &cobra.Command{
	Use:     "rc {name|id} newName",
	Aliases: []string{"RenameCos"},
	Short:   "Rename Cos",
	Long:    `Rename a class of service (COS), given by name or zimbraId. Its zimbraId is kept, so the accounts using it are not affected.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.RenameCos(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(rcCmd)
}
