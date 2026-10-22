package cmd

import (
	"fmt"
	"rprov/soap/admin"
	"strconv"

	"github.com/spf13/cobra"
)

var saVerbose bool

var saCmd = &cobra.Command{
	Use:     "sa [-v] {ldap-query} [limit {limit}] [offset {offset}] [sortBy {attr}] [sortAscending 0|1] [domain {domain}]",
	Aliases: []string{"SearchAccounts"},
	Short:   "Search Accounts",
	Long: `List the accounts matching an LDAP query, e.g. "(zimbraAccountStatus=locked)".
Without a limit, all matching accounts are listed. Results are sorted by name in ascending order by default.
With -v, display all the attributes of each account.`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires an LDAP query")
		}

		if len(args)%2 != 1 {
			return fmt.Errorf("option %q has no value", args[len(args)-1])
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		opts := admin.SearchOptions{
			Query:         args[0],
			SortBy:        "name",
			SortAscending: true,
			Verbose:       saVerbose,
		}

		for i := 1; i < len(args); i += 2 {
			name, value := args[i], args[i+1]

			switch name {
			case "limit", "offset":
				n, err := strconv.Atoi(value)
				if err != nil || n < 0 {
					return fmt.Errorf("invalid %s %q: expected a non-negative integer", name, value)
				}

				if name == "limit" {
					opts.Limit = n
				} else {
					opts.Offset = n
				}
			case "sortBy":
				if value == "" {
					return fmt.Errorf("empty sortBy attribute")
				}

				opts.SortBy = value
			case "sortAscending":
				switch value {
				case "0":
					opts.SortAscending = false
				case "1":
					opts.SortAscending = true
				default:
					return fmt.Errorf("invalid sortAscending %q: expected 0 or 1", value)
				}
			case "domain":
				if value == "" {
					return fmt.Errorf("empty domain")
				}

				opts.Domain = value
			default:
				return fmt.Errorf("unknown option %q", name)
			}
		}

		conf, err := loadConfig()
		if err != nil {
			return err
		}

		return admin.SearchAccounts(conf, opts)
	},
}

func init() {
	rootCmd.AddCommand(saCmd)

	saCmd.Flags().BoolVarP(&saVerbose, "verbose", "v", false, "Display the attributes of each account")
}
