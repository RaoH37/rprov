package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

const searchLimit = 2000

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
