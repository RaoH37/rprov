package admin

import "fmt"

func CreateDomain(conf map[string]string, name string, attributes []string) error {
	if err := validateDomainName(name); err != nil {
		return err
	}

	attrs, err := buildAttrs(attributes, false)

	if err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	response, err := zcs.CreateDomainRequest(name, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Domain) == 0 {
		return fmt.Errorf("server returned no domain after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Domain[0].ID)

	return nil
}
