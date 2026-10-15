package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func CreateAccount(conf map[string]string, name string, password string, attributes []string) error {
	attrs, err := buildAttrs(attributes)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	response, err := zcs.CreateAccountRequest(name, password, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Account) == 0 {
		return fmt.Errorf("server returned no account after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Account[0].ID)

	return nil
}

func buildAttrs(args []string) (zimbraCommon.AttrsNode, error) {
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("attribute %q has no value", args[len(args)-1])
	}

	attrs := make(zimbraCommon.AttrsNode, 0, len(args)/2)

	for i := 0; i < len(args); i += 2 {
		name := args[i]

		if len(name) > 0 && name[0] == '+' {
			name = name[1:]
		}

		if name == "" {
			return nil, fmt.Errorf("empty attribute name")
		}

		attrs = append(attrs, zimbraCommon.AttrNode{
			Name:  name,
			Value: args[i+1],
		})
	}

	return attrs, nil
}
