package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gasCmd = &cobra.Command{
	Use:     "gas [service]",
	Aliases: []string{"GetAllServers"},
	Short:   "Get All Servers",
	Long:    `List all servers. If a service is given (e.g. mailbox, mta, proxy), list only the servers running that service.`,

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		service := ""
		if len(args) > 0 {
			service = args[0]
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetAllServers(conf, service)
	},
}

func init() {
	rootCmd.AddCommand(gasCmd)
}
