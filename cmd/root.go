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

var attributes []string

var batchFile string

var localConfig map[string]string

var rootCmd = &cobra.Command{
	Use:   "rprov",
	Short: "zmprov clone written in Go",
	Long: `zmprov clone written in Go.
With -f, run the commands of a file, one per line, with a single authentication.
Empty lines and lines starting with # are ignored. Quote arguments with ' or ", and use "" for an empty value.
A failed line is reported with its line number and the next lines are still run.`,

	Args: cobra.NoArgs,

	SilenceErrors: true,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if batchFile != "" {
			return runBatch(batchFile)
		}

		fmt.Println("Use 'rprov help' to see the available commands.")

		return nil
	}

	rootCmd.Flags().StringVarP(&batchFile, "file", "f", "", "Run the commands of a file, one per line")

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
	if localConfig == nil {
		path, err := config.FindLocalConfig()
		if err != nil {
			return nil, err
		}

		parsed, err := config.ParseLocalConfig(path)
		if err != nil {
			return nil, err
		}

		localConfig = parsed
	}

	conf := make(map[string]string, len(localConfig)+3)
	for k, v := range localConfig {
		conf[k] = v
	}

	conf["zimbra_admin_service_url"] = serverURL
	conf["zimbra_admin_auth_login"] = authLogin
	conf["zimbra_admin_auth_password"] = authPassword

	return conf, nil
}
