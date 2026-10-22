package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var cpcCmd = &cobra.Command{
	Use:     "cpc {src-cos|id} dest-cos",
	Aliases: []string{"CopyCos"},
	Short:   "Copy Cos",
	Long:    `Create a class of service (COS) with the attributes of an existing one, given by name or zimbraId, and print its zimbraId.`,

	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.CopyCos(conf, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(cpcCmd)
}
