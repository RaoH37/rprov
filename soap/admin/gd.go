package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetDomain(conf map[string]string, key string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	attrs := zimbraCommon.StringList(attributes)

	id, name := splitIdOrName(key)

	response, err := zcs.GetDomainRequest(id, name, attrs, true)

	if err != nil {
		return err
	}

	if len(response.Content.Domains) == 0 {
		return fmt.Errorf("no such domain: %s", key)
	}

	domain := response.Content.Domains[0]

	fmt.Printf("# name %v\n", domain.Name)

	sort.Slice(domain.Attrs, func(i, j int) bool {
		return domain.Attrs[i].Name < domain.Attrs[j].Name
	})

	for i := 0; i < len(domain.Attrs); i++ {
		attr := domain.Attrs[i]
		outputAttr(attr)

		if attr.Name == "zimbraDomainDefaultCOSId" {
			if cosName, cosErr := getCosName(zcs, attr.Value); cosErr == nil {
				fmt.Printf("zimbraDomainDefaultCOSName: %v\n", cosName)
			}
		}
	}

	return nil
}
