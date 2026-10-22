package admin

import (
	"errors"

	"github.com/RaoH37/gosoap/zadmin"
)

type session struct {
	url      string
	login    string
	password string
	zcs      *zadmin.ZAdmin
}

var current *session

type AuthError struct {
	Err error
}

func (e *AuthError) Error() string {
	return "authentication failed: " + e.Err.Error()
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

func Authenticate(conf map[string]string) (*zadmin.ZAdmin, error) {
	url := conf["zimbra_admin_service_url"]
	login := conf["zimbra_admin_auth_login"]
	password := getPassword(conf)

	if current != nil && current.url == url && current.login == login && current.password == password {
		return current.zcs, nil
	}

	zcs := zadmin.NewZAdmin(url, login, password, zadmin.WithInsecureMode())

	resp, err := zcs.AuthRequest()

	if err != nil {
		return nil, &AuthError{Err: err}
	}

	if len(resp.Content.TOKEN) == 0 {
		return nil, &AuthError{Err: errors.New("response contains no token")}
	}

	zcs.SetToken(resp.Content.TOKEN[0].Content)

	current = &session{url: url, login: login, password: password, zcs: zcs}

	return zcs, nil
}

func getPassword(conf map[string]string) string {
	if conf["zimbra_admin_auth_password"] != "" {
		return conf["zimbra_admin_auth_password"]
	}

	return conf["zimbra_ldap_password"]
}
