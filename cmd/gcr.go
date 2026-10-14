package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gcrCmd = &cobra.Command{
	Use:     "gcr name",
	Aliases: []string{"GetCalendarResource"},
	Short:   "Get Calendar Resource Attributes",
	Long:    `todo`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetCalendarResource(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gcrCmd)
}
