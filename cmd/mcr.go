package cmd

import (
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var mcrCmd = &cobra.Command{
	Use:     "mcr {name@domain|id} attr1 value1 [attr2 value2...]",
	Aliases: []string{"ModifyCalendarResource"},
	Short:   "Modify Calendar Resource",
	Long: `Modify the attributes of a calendar resource, given by name or zimbraId.
Prefix an attribute with + to add a value or - to remove a value of a multi-valued attribute.
Without a prefix, the given values replace the current ones; repeat an attribute to set several values.
An empty value ("") removes the attribute.`,

	Args: cobra.MinimumNArgs(3),

	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		args = args[1:]

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.ModifyCalendarResource(conf, key, args)
	},
}

func init() {
	rootCmd.AddCommand(mcrCmd)
}
