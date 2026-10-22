package admin

func RemoveAccountAlias(conf map[string]string, key string, alias string) error {
	if err := validateAddress(alias); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getAccountId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.RemoveAccountAliasRequest(id, alias)
}
