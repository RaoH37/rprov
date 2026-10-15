package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gadlCmd = &cobra.Command{
	Use:     "gadl [domain]",
	Aliases: []string{"GetAllDistributionLists"},
	Short:   "Get All Distribution Lists",
	Long:    `List all distribution lists.`,

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

		return admin.GetAllDistributionLists(conf, domain, attributes)
	},
}

func init() {
	rootCmd.AddCommand(gadlCmd)

	gadlCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
