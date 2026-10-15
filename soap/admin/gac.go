package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAllCos(conf map[string]string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeCoses,
	}

	attrs := zimbraCommon.StringList(attributes)

	offset := 0

	for {
		response, err := zcs.SearchDirectoryRequest("", 1_000_000, searchLimit, offset, "", false, false, "cn", types, true, attrs, false)

		if err != nil {
			return err
		}

		for i := 0; i < len(response.Content.Coses); i++ {
			fmt.Printf("%v\n", response.Content.Coses[i].Name)
		}

		if len(response.Content.Coses) < searchLimit {
			break
		}

		offset += searchLimit
	}

	return nil
}
