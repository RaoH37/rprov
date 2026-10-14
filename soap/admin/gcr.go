package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetCalendarResource(conf map[string]string, name string, attributes []string) {
	zcs := Authenticate(conf)

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetCalendarResourceRequest("", name, attrs, true)

	if err != nil {
		panic(err)
	}

	resource := response.Content.CalResource[0]

	fmt.Printf("# name %v\n", resource.Name)

	sort.Slice(resource.Attrs, func(i, j int) bool {
		return resource.Attrs[i].Name < resource.Attrs[j].Name
	})

	for i := 0; i < len(resource.Attrs); i++ {
		attr := resource.Attrs[i]
		outputAttr(attr)

		if attr.Name == "zimbraCOSId" {
			if cosName, cosErr := getCosName(zcs, attr.Value); cosErr == nil {
				fmt.Printf("zimbraCOSName: %v\n", cosName)
			}
		}
	}
}
