package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func ModifyAccount(conf map[string]string, key string, attributes []string) error {
	attrs, err := buildAttrs(attributes, true)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(key)

	response, err := zcs.GetAccountRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return err
	}

	if len(response.Content.Account) == 0 || response.Content.Account[0].ID == "" {
		return fmt.Errorf("no such account: %s", key)
	}

	return zcs.ModifyAccountRequest(response.Content.Account[0].ID, attrs)
}
