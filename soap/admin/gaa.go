package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAllAccounts(conf map[string]string, domain string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeAccounts,
	}

	attrs := zimbraCommon.StringList(attributes)

	offset := 0

	for {
		response, err := zcs.SearchDirectoryRequest("", 1_000_000, searchLimit, offset, domain, false, false, "mail", types, true, attrs, false)

		if err != nil {
			return err
		}

		for i := 0; i < len(response.Content.Accounts); i++ {
			fmt.Printf("%v\n", response.Content.Accounts[i].Name)
		}

		if len(response.Content.Accounts) < searchLimit {
			break
		}

		offset += searchLimit
	}

	return nil
}
