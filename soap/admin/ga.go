package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetAccount(conf map[string]string, name string, attributes []string) {
	zcs := Authenticate(conf)

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetAccountRequest("", name, attrs, true)

	if err != nil {
		panic(err)
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
}
