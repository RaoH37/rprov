package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAllDistributionLists(conf map[string]string, domain string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeDistributionLists,
	}

	attrs := zimbraCommon.StringList(attributes)

	offset := 0

	for {
		response, err := zcs.SearchDirectoryRequest("", 1_000_000, searchLimit, offset, domain, false, false, "mail", types, true, attrs, false)

		if err != nil {
			return err
		}

		for i := 0; i < len(response.Content.Dls); i++ {
			fmt.Printf("%v\n", response.Content.Dls[i].Name)
		}

		if len(response.Content.Dls) < searchLimit {
			break
		}

		offset += searchLimit
	}

	return nil
}
