package admin

func DeleteDomain(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getDomainId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.DeleteDomainRequest(id)
}
