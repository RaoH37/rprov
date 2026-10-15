package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetCos(conf map[string]string, name string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetCosRequest("", name, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Cos) == 0 {
		return fmt.Errorf("no such cos: %s", name)
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

	return nil
}
