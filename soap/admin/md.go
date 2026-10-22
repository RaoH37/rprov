package admin

func ModifyDomain(conf map[string]string, key string, attributes []string) error {
	attrs, err := buildAttrs(attributes, true)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getDomainId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.ModifyDomainRequest(id, attrs)
}
