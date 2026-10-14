package admin

import "fmt"

func GetAllServers(conf map[string]string, service string) {
	zcs := Authenticate(conf)

	response, err := zcs.GetAllServersRequest(service)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(response.Content.Servers); i++ {
		fmt.Printf("%v\n", response.Content.Servers[i].Name)
	}
}
