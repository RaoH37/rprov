package admin

import "fmt"

func GetAllServers(conf map[string]string, service string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	response, err := zcs.GetAllServersRequest(service)

	if err != nil {
		return err
	}

	for i := 0; i < len(response.Content.Servers); i++ {
		fmt.Printf("%v\n", response.Content.Servers[i].Name)
	}

	return nil
}
