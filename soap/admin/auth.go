package admin

import (
	"errors"

	"github.com/RaoH37/gosoap/zadmin"
)

func Authenticate(conf map[string]string) (*zadmin.ZAdmin, error) {
	zcs := zadmin.NewZAdmin(
		conf["zimbra_admin_service_url"],
		conf["zimbra_admin_auth_login"],
		getPassword(conf),
		zadmin.WithInsecureMode(),
	)

	resp, err := zcs.AuthRequest()

	if err != nil {
		return nil, err
	}

	if len(resp.Content.TOKEN) == 0 {
		return nil, errors.New("authentication response contains no token")
	}

	zcs.SetToken(resp.Content.TOKEN[0].Content)

	return zcs, nil
}

func getPassword(conf map[string]string) string {
	if conf["zimbra_admin_auth_password"] != "" {
		return conf["zimbra_admin_auth_password"]
	}

	return conf["zimbra_ldap_password"]
}
