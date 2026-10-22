package admin

func RenameCalendarResource(conf map[string]string, key string, newName string) error {
	if err := validateAddress(newName); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getCalendarResourceId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.RenameCalendarResourceRequest(id, newName)
}
