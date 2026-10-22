package admin

func RenameDistributionList(conf map[string]string, key string, newName string) error {
	if err := validateAddress(newName); err != nil {
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

	return zcs.RenameDistributionListRequest(id, newName)
}
