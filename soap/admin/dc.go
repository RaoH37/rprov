package admin

func DeleteCos(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getCosId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.DeleteCosRequest(id)
}
