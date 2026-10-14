package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetCos(conf map[string]string, name string, attributes []string) {
	zcs := Authenticate(conf)

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetCosRequest("", name, attrs)

	if err != nil {
		panic(err)
	}

	cos := response.Content.Cos[0]

	fmt.Printf("# name %v\n", cos.Name)

	sort.Slice(cos.Attrs, func(i, j int) bool {
		return cos.Attrs[i].Name < cos.Attrs[j].Name
	})

	for i := 0; i < len(cos.Attrs); i++ {
		attr := cos.Attrs[i]
		outputAttr(attr)
	}
}
