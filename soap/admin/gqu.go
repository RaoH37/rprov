package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func GetQuotaUsage(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(key)

	servers, err := zcs.GetServerRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return err
	}

	if len(servers.Content.Server) == 0 || servers.Content.Server[0].ID == "" {
		return fmt.Errorf("no such server: %s", key)
	}

	response, err := zcs.GetQuotaUsageRequest(servers.Content.Server[0].ID, "", false)

	if err != nil {
		return err
	}

	for i := 0; i < len(response.Content.Account); i++ {
		account := response.Content.Account[i]
		fmt.Printf("%v %v %v\n", account.Name, account.Limit, account.Used)
	}

	return nil
}
