package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gacCmd = &cobra.Command{
	Use:     "gac",
	Aliases: []string{"GetAllCos"},
	Short:   "Get All Cos",
	Long:    `List all COS.`,

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(attributes) == 0 {
			attributes = []string{"mail"}
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAllCos(conf, attributes)
	},
}

func init() {
	rootCmd.AddCommand(gacCmd)

	gacCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
