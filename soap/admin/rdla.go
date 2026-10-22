package admin

func RemoveDistributionListAlias(conf map[string]string, key string, alias string) error {
	if err := validateAddress(alias); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getDistributionListId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.RemoveDistributionListAliasRequest(id, alias)
}
