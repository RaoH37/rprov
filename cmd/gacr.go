package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gacrCmd = &cobra.Command{
	Use:     "gacr [domain]",
	Aliases: []string{"GetAllCalendarResources"},
	Short:   "Get All Calendar Resources",
	Long:    `List all calendar resources.`,

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

		return admin.GetAllCalendarResources(conf, domain, attributes)
	},
}

func init() {
	rootCmd.AddCommand(gacrCmd)

	gacrCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
