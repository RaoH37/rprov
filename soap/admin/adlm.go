package admin

func AddDistributionListMember(conf map[string]string, key string, members []string) error {
	if err := validateAddresses(members); err != nil {
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

	return zcs.AddDistributionListMemberRequest(id, members)
}
