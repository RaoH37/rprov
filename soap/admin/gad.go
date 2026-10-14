package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAllDomains(conf map[string]string, attributes []string) {
	zcs := Authenticate(conf)

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeDomains,
	}

	attrs := zimbraCommon.StringList(attributes)

	offset := 0
	//limit := 500

	for {
		response, err := zcs.SearchDirectoryRequest("", 1_000_000, searchLimit, offset, "", false, false, "zimbraDomainName", types, true, attrs, false)

		if err != nil {
			panic(err)
		}

		for i := 0; i < len(response.Content.Domains); i++ {
			fmt.Printf("%v\n", response.Content.Domains[i].Name)
		}

		if len(response.Content.Domains) < searchLimit {
			break
		}

		offset += searchLimit
	}
}
