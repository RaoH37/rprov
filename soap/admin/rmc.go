package admin

import (
	"fmt"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func RecalculateMailboxCounts(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(key)

	accounts, err := zcs.GetAccountRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return err
	}

	if len(accounts.Content.Account) == 0 || accounts.Content.Account[0].ID == "" {
		return fmt.Errorf("no such account: %s", key)
	}

	account := accounts.Content.Account[0]

	response, err := zcs.RecalculateMailboxCountsRequest(account.ID)

	if err != nil {
		return err
	}

	if len(response.Content.Mbox) == 0 {
		return fmt.Errorf("server returned no mailbox for %s", account.Name)
	}

	fmt.Printf("account: %v\n", account.Name)
	fmt.Printf("mailbox:\n")
	fmt.Printf("    quota used: %v\n", response.Content.Mbox[0].Used)

	return nil
}
