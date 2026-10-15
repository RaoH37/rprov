package admin

import (
	"fmt"
	"sort"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetServer(conf map[string]string, name string, attributes []string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	attrs := zimbraCommon.StringList(attributes)

	response, err := zcs.GetServerRequest("", name, attrs, true)

	if err != nil {
		return err
	}

	if len(response.Content.Server) == 0 {
		return fmt.Errorf("no such server: %s", name)
	}

	server := response.Content.Server[0]

	fmt.Printf("# name %v\n", server.Name)

	sort.Slice(server.Attrs, func(i, j int) bool {
		return server.Attrs[i].Name < server.Attrs[j].Name
	})

	for i := 0; i < len(server.Attrs); i++ {
		attr := server.Attrs[i]
		outputAttr(attr)
	}

	return nil
}
