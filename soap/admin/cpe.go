package admin

import (
	"fmt"
	"strings"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func ChangePrimaryEmail(conf map[string]string, key string, newName string) error {
	if err := validateAddress(newName); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(key)

	response, err := zcs.GetAccountRequest(id, name, zimbraCommon.StringList{"zimbraId", "zimbraMailAlias"}, false)

	if err != nil {
		return err
	}

	if len(response.Content.Account) == 0 || response.Content.Account[0].ID == "" {
		return fmt.Errorf("no such account: %s", key)
	}

	account := response.Content.Account[0]

	if strings.EqualFold(account.Name, newName) {
		return fmt.Errorf("%s is already the primary email of the account", newName)
	}

	isAlias := false

	for _, attr := range account.Attrs {
		if attr.Name == "zimbraMailAlias" && strings.EqualFold(attr.Value, newName) {
			isAlias = true
			break
		}
	}

	if isAlias {
		if err := zcs.RemoveAccountAliasRequest(account.ID, newName); err != nil {
			return fmt.Errorf("cannot remove alias %s before renaming: %w", newName, err)
		}
	}

	if err := zcs.RenameAccountRequest(account.ID, newName); err != nil {
		if isAlias {
			if restoreErr := zcs.AddAccountAliasRequest(account.ID, newName); restoreErr != nil {
				return fmt.Errorf("cannot rename %s to %s: %w (alias %s was removed and could not be restored: %v)", account.Name, newName, err, newName, restoreErr)
			}
		}

		return fmt.Errorf("cannot rename %s to %s: %w", account.Name, newName, err)
	}

	if err := zcs.AddAccountAliasRequest(account.ID, account.Name); err != nil {
		return fmt.Errorf("account renamed to %s, but %s could not be kept as an alias: %w", newName, account.Name, err)
	}

	return nil
}
