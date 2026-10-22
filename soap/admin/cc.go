package admin

import "fmt"

func CreateCos(conf map[string]string, name string, attributes []string) error {
	if err := validateCosName(name); err != nil {
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

	response, err := zcs.CreateCosRequest(name, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Cos) == 0 {
		return fmt.Errorf("server returned no cos after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Cos[0].ID)

	return nil
}
