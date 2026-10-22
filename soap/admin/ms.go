package admin

func ModifyServer(conf map[string]string, key string, attributes []string) error {
	attrs, err := buildAttrs(attributes, true)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getServerId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.ModifyServerRequest(id, attrs)
}
