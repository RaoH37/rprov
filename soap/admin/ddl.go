package admin

func DeleteDistributionList(conf map[string]string, key string, cascadeDelete bool) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getDistributionListId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.DeleteDistributionListRequest(id, cascadeDelete)
}
