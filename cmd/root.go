package cmd

import (
	"fmt"
	"os"
	"rprov/internal/config"

	"github.com/spf13/cobra"
)

var serverURL string
var authLogin string
var authPassword string

// Variables holding the flag values
var attributes []string

var rootCmd = &cobra.Command{
	Use:   "rprov",
	Short: "zmprov clone written in Go",
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is given
		fmt.Println("Use 'rprov help' to see the available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&serverURL,
		"url",
		"",
		"https://localhost:7071/service/admin/soap",
		"Zimbra Admin SOAP API URL",
	)

	rootCmd.PersistentFlags().StringVarP(
		&authLogin,
		"login",
		"",
		"zimbra",
		"Authentication account login",
	)

	rootCmd.PersistentFlags().StringVarP(
		&authPassword,
		"password",
		"",
		"",
		"Authentication account password",
	)
}

func loadConfig() (map[string]string, error) {
	path, err := config.FindLocalConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	conf, err := config.ParseLocalConfig(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	conf["zimbra_admin_service_url"] = serverURL
	conf["zimbra_admin_auth_login"] = authLogin
	conf["zimbra_admin_auth_password"] = authPassword

	return conf, nil
}
