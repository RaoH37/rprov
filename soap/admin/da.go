package admin

func DeleteAccount(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getAccountId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.DeleteAccountRequest(id)
}
