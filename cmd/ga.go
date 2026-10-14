package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gaCmd = &cobra.Command{
	Use:     "ga name",
	Aliases: []string{"GetAccount"},
	Short:   "Get Account Attributes",
	Long:    `todo`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetAccount(conf, name, args)
	},
}

func init() {
	rootCmd.AddCommand(gaCmd)
}
