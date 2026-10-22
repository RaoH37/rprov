package admin

import "fmt"

func CreateDistributionList(conf map[string]string, name string, dynamic bool, attributes []string) error {
	if err := validateAddress(name); err != nil {
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

	response, err := zcs.CreateDistributionListRequest(name, dynamic, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Dl) == 0 {
		return fmt.Errorf("server returned no distribution list after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Dl[0].ID)

	return nil
}
