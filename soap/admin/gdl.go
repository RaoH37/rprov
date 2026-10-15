package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetDistributionList(conf map[string]string, name string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetDistributionListRequest("", name, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Dl) == 0 {
		return fmt.Errorf("no such distribution list: %s", name)
	}

	dl := response.Content.Dl[0]

	fmt.Printf("# name %v\n", dl.Name)

	sort.Slice(dl.Attrs, func(i, j int) bool {
		return dl.Attrs[i].Name < dl.Attrs[j].Name
	})

	for i := 0; i < len(dl.Attrs); i++ {
		attr := dl.Attrs[i]
		outputAttr(attr)
	}

	return nil
}
