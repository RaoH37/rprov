package admin

import "fmt"

func CreateAccount(conf map[string]string, name string, password string, attributes []string) error {
	attrs, err := buildAttrs(attributes, false)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	response, err := zcs.CreateAccountRequest(name, password, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Account) == 0 {
		return fmt.Errorf("server returned no account after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Account[0].ID)

	return nil
}
