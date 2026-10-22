package admin

import (
	"errors"
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

type SearchOptions struct {
	Query         string
	Domain        string
	SortBy        string
	SortAscending bool
	// Limit is the maximum number of accounts to return; 0 returns them all.
	Limit      int
	Offset     int
	Verbose    bool
	Attributes []string
}

func SearchAccounts(conf map[string]string, opts SearchOptions) error {
	if opts.Query == "" {
		return errors.New("empty LDAP query")
	}

	if opts.Limit < 0 || opts.Offset < 0 {
		return errors.New("limit and offset must not be negative")
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	return searchAccounts(zcs, opts)
}

func searchAccounts(zcs *zadmin.ZAdmin, opts SearchOptions) error {
	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeAccounts,
	}

	attrs := zimbraCommon.StringList(opts.Attributes)

	if !opts.Verbose && len(attrs) == 0 {
		attrs = zimbraCommon.StringList{"mail"}
	}

	offset := opts.Offset
	remaining := opts.Limit
	first := true

	for {
		pageSize := searchLimit

		if opts.Limit > 0 && remaining < pageSize {
			pageSize = remaining
		}

		response, err := zcs.SearchDirectoryRequest(opts.Query, 1_000_000, pageSize, offset, opts.Domain, false, false, opts.SortBy, types, opts.SortAscending, attrs, false)

		if err != nil {
			return err
		}

		for i := 0; i < len(response.Content.Accounts); i++ {
			account := response.Content.Accounts[i]

			if !opts.Verbose {
				fmt.Printf("%v\n", account.Name)
				continue
			}

			if !first {
				fmt.Println()
			}

			first = false

			fmt.Printf("# name %v\n", account.Name)

			sort.Slice(account.Attrs, func(i, j int) bool {
				return account.Attrs[i].Name < account.Attrs[j].Name
			})

			for j := 0; j < len(account.Attrs); j++ {
				outputAttr(account.Attrs[j])
			}
		}

		if len(response.Content.Accounts) < pageSize {
			break
		}

		offset += pageSize

		if opts.Limit > 0 {
			remaining -= pageSize

			if remaining == 0 {
				break
			}
		}
	}

	return nil
}
