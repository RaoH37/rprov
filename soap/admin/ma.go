package admin

func ModifyAccount(conf map[string]string, key string, attributes []string) error {
	attrs, err := buildAttrs(attributes, true)

	if err != nil {
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

	return zcs.ModifyAccountRequest(id, attrs)
}
