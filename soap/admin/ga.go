package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAccount(conf map[string]string, key string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	attrs := zimbraCommon.StringList(attributes)

	id, name := splitIdOrName(key)

	response, err := zcs.GetAccountRequest(id, name, attrs, true)

	if err != nil {
		return err
	}

	if len(response.Content.Account) == 0 {
		return fmt.Errorf("no such account: %s", key)
	}

	account := response.Content.Account[0]

	fmt.Printf("# name %v\n", account.Name)

	sort.Slice(account.Attrs, func(i, j int) bool {
		return account.Attrs[i].Name < account.Attrs[j].Name
	})

	for i := 0; i < len(account.Attrs); i++ {
		attr := account.Attrs[i]
		outputAttr(attr)

		if attr.Name == "zimbraCOSId" {
			if cosName, cosErr := getCosName(zcs, attr.Value); cosErr == nil {
				fmt.Printf("zimbraCOSName: %v\n", cosName)
			}
		}
	}

	return nil
}
