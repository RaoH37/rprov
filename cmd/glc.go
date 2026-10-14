package cmd

import (
	"fmt"
	"rprov/internal/config"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var showSecrets bool

var glcCmd = &cobra.Command{
	Use:     "glc",
	Aliases: []string{"getLocalConfig"},
	Short:   "Display the values read from localconfig.xml",
	Long:    `todo`,

	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.FindLocalConfig()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		conf, err := config.ParseLocalConfig(path)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		keys := make([]string, 0, len(conf))
		for k := range conf {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := conf[k]
			if !showSecrets && strings.HasSuffix(k, "_password") {
				v = "********"
			}
			fmt.Printf("%s = %s\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(glcCmd)

	glcCmd.Flags().BoolVarP(&showSecrets, "secrets", "s", false, "Display passwords in clear text")
}
