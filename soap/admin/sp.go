package admin

import "errors"

func SetPassword(conf map[string]string, key string, password string) error {
	if password == "" {
		return errors.New("password must not be empty")
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getAccountId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.SetPasswordRequest(id, password)
}
