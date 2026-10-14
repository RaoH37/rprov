package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gdlCmd = &cobra.Command{
	Use:     "gdl name",
	Aliases: []string{"GetDistributionList"},
	Short:   "Get DistributionList Attributes",
	Long:    `todo`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetDistributionList(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gdlCmd)
}
