package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetDistributionList(conf map[string]string, name string, attributes []string) {
	zcs := Authenticate(conf)

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetDistributionListRequest("", name, attrs)

	if err != nil {
		panic(err)
	}

	dl := response.Content.Dl[0]

	fmt.Printf("# name %v\n", dl.Name)

	sort.Slice(dl.Attrs, func(i, j int) bool {
		return dl.Attrs[i].Name < dl.Attrs[j].Name
	})

	for i := 0; i < len(dl.Attrs); i++ {
		attr := dl.Attrs[i]
		outputAttr(attr)

		//if attr.Name == "zimbraCOSId" {
		//	if cosName, cosErr := getCosName(zcs, attr.Value); cosErr == nil {
		//		fmt.Printf("zimbraCOSName: %v\n", cosName)
		//	}
		//}
	}
}
