package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:     "gc {name|id} [attr1 [attr2...]]",
	Aliases: []string{"GetCos"},
	Short:   "Get Cos Attributes",
	Long:    `Display the attributes of a class of service (COS), given by name or zimbraId, sorted by name. If attributes are given, display only those.`,

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		args = append(args[:0], args[1:]...)

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.GetCos(conf, key, args)
	},
}

func init() {
	rootCmd.AddCommand(gcCmd)
}
