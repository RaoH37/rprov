package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gasCmd = &cobra.Command{
	Use:     "gas", // [domain] in brackets means optional
	Aliases: []string{"GetAllServers"},
	Short:   "Get All Servers",
	Long:    `todo`,

	// Positional argument validation
	// At most 1 argument (the domain)
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		service := ""
		if len(args) > 0 {
			service = args[0]
		}

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetAllServers(conf, service)
	},
}

func init() {
	rootCmd.AddCommand(gasCmd)
}
