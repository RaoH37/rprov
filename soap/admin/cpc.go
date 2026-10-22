package admin

import "fmt"

func CopyCos(conf map[string]string, key string, newName string) error {
	if err := validateCosName(newName); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(key)

	response, err := zcs.CopyCosRequest(id, name, newName)

	if err != nil {
		return err
	}

	if len(response.Content.Cos) > 0 && response.Content.Cos[0].ID != "" {
		fmt.Printf("%v\n", response.Content.Cos[0].ID)
		return nil
	}

	newId, err := getCosId(zcs, newName)

	if err != nil {
		return fmt.Errorf("cos copied to %s, but its zimbraId could not be read: %w", newName, err)
	}

	fmt.Printf("%v\n", newId)

	return nil
}
