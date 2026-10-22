package admin

import "fmt"

func GetQuotaUsage(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	serverId, err := getServerId(zcs, key)

	if err != nil {
		return err
	}

	response, err := zcs.GetQuotaUsageRequest(serverId, "", false)

	if err != nil {
		return err
	}

	for i := 0; i < len(response.Content.Account); i++ {
		account := response.Content.Account[i]
		fmt.Printf("%v %v %v\n", account.Name, account.Limit, account.Used)
	}

	return nil
}
