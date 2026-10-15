package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gadCmd = &cobra.Command{
	Use:     "gad",
	Aliases: []string{"GetAllDomains"},
	Short:   "Get All Domains",
	Long:    `List all domains.`,

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(attributes) == 0 {
			attributes = []string{"zimbraDomainName"}
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAllDomains(conf, attributes)
	},
}

func init() {
	rootCmd.AddCommand(gadCmd)

	gadCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
