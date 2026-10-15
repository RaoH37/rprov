package admin

import (
	"fmt"
	"regexp"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

const searchLimit = 2000

var uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func outputAttr(attr zimbraCommon.AttrNode) {
	fmt.Printf("%v: %v\n", attr.Name, attr.Value)
}

func getCosName(zcs *zadmin.ZAdmin, id string) (string, error) {
	response, err := zcs.GetCosRequest(id, "", zimbraCommon.StringList{"cn"})

	if err != nil {
		return "", err
	}

	if len(response.Content.Cos) == 0 {
		return "", fmt.Errorf("no such cos with id %s", id)
	}

	return response.Content.Cos[0].Name, nil
}

func splitIdOrName(key string) (string, string) {
	if uuidPattern.MatchString(key) {
		return key, ""
	}

	return "", key
}

func buildAttrs(args []string, keepModifiers bool) (zimbraCommon.AttrsNode, error) {
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("attribute %q has no value", args[len(args)-1])
	}

	attrs := make(zimbraCommon.AttrsNode, 0, len(args)/2)

	for i := 0; i < len(args); i += 2 {
		name := args[i]
		base := name

		if len(name) > 0 && (name[0] == '+' || (keepModifiers && name[0] == '-')) {
			base = name[1:]

			if !keepModifiers {
				name = base
			}
		}

		if base == "" {
			return nil, fmt.Errorf("empty attribute name")
		}

		attrs = append(attrs, zimbraCommon.AttrNode{
			Name:  name,
			Value: args[i+1],
		})
	}

	return attrs, nil
}
