package cmd

import (
	"fmt"
	"rprov/soap/admin"

	"github.com/spf13/cobra"
)

var ddlCmd = &cobra.Command{
	Use:     "ddl {list@domain|id} [true|false]",
	Aliases: []string{"DeleteDistributionList"},
	Short:   "Delete Distribution List",
	Long: `Delete a distribution list, given by name or zimbraId.
The optional true|false argument is sent as the cascadeDelete flag of the request (false by default).
There is no confirmation prompt and the deletion cannot be undone.`,

	Args: cobra.RangeArgs(1, 2),

	RunE: func(cmd *cobra.Command, args []string) error {
		cascadeDelete := false

		if len(args) == 2 {
			switch args[1] {
			case "true":
				cascadeDelete = true
			case "false":
			default:
				return fmt.Errorf("invalid cascade delete value %q: expected true or false", args[1])
			}
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.DeleteDistributionList(conf, args[0], cascadeDelete)
	},
}

func init() {
	rootCmd.AddCommand(ddlCmd)
}
