package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var gacCmd = &cobra.Command{
	Use:     "gac", // [domain] in brackets means optional
	Aliases: []string{"GetAllCos"},
	Short:   "Get All Cos",
	Long:    `List all COS.`,

	// Positional argument validation
	// At most 1 argument (the domain)
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		//domain := ""
		//if len(args) > 0 {
		//	domain = args[0]
		//}

		if len(attributes) == 0 {
			attributes = []string{"mail"}
		}

		conf, err := loadConfig()
		if err != nil {
			return
		}

		admin.GetAllCos(conf, attributes)
	},
}

func init() {
	// Register the command on the root
	rootCmd.AddCommand(gacCmd)

	// Flag definitions for 'gaa'
	// StringSliceVarP allows passing the same flag several times: --attr sn --attr cn
	// Or comma-separated: --attr sn,cn
	// "attributes": long name
	// "a": shorthand -a
	// []string{}: default value (empty)
	// "description"
	gacCmd.Flags().StringSliceVarP(&attributes, "attr", "a", []string{}, "Attributes to display")
}
