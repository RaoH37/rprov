package admin

func RenameCos(conf map[string]string, key string, newName string) error {
	if err := validateCosName(newName); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getCosId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.RenameCosRequest(id, newName)
}
