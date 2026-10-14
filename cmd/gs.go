package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gsCmd = &cobra.Command{
	Use:     "gs name",
	Aliases: []string{"GetServer"},
	Short:   "Get Server Attributes",
	Long:    `todo`,

	//Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetServer(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gsCmd)
}
