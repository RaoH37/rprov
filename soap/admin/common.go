package admin

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

const searchLimit = 2000

var uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func outputAttr(attr zimbraCommon.AttrNode) {
	fmt.Printf("%v: %v\n", attr.Name, attr.Value)
}

func getCosName(zcs *zadmin.ZAdmin, id string) (string, error) {
	response, err := zcs.GetCosRequest(id, "", zimbraCommon.StringList{"cn"})

	if err != nil {
		return "", err
	}

	if len(response.Content.Cos) == 0 {
		return "", fmt.Errorf("no such cos with id %s", id)
	}

	return response.Content.Cos[0].Name, nil
}

func splitIdOrName(key string) (string, string) {
	if uuidPattern.MatchString(key) {
		return key, ""
	}

	return "", key
}

func getAccountId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetAccountRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return "", err
	}

	if len(response.Content.Account) == 0 || response.Content.Account[0].ID == "" {
		return "", fmt.Errorf("no such account: %s", key)
	}

	return response.Content.Account[0].ID, nil
}

func getDomainId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetDomainRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return "", err
	}

	if len(response.Content.Domains) == 0 || response.Content.Domains[0].ID == "" {
		return "", fmt.Errorf("no such domain: %s", key)
	}

	return response.Content.Domains[0].ID, nil
}

func getDistributionListId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetDistributionListRequest(id, name, zimbraCommon.StringList{"zimbraId"})

	if err != nil {
		return "", err
	}

	if len(response.Content.Dl) == 0 || response.Content.Dl[0].ID == "" {
		return "", fmt.Errorf("no such distribution list: %s", key)
	}

	return response.Content.Dl[0].ID, nil
}

func getCalendarResourceId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetCalendarResourceRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return "", err
	}

	if len(response.Content.CalResource) == 0 || response.Content.CalResource[0].ID == "" {
		return "", fmt.Errorf("no such calendar resource: %s", key)
	}

	return response.Content.CalResource[0].ID, nil
}

func getServerId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetServerRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return "", err
	}

	if len(response.Content.Server) == 0 || response.Content.Server[0].ID == "" {
		return "", fmt.Errorf("no such server: %s", key)
	}

	return response.Content.Server[0].ID, nil
}

func getCosId(zcs *zadmin.ZAdmin, key string) (string, error) {
	id, name := splitIdOrName(key)

	response, err := zcs.GetCosRequest(id, name, zimbraCommon.StringList{"zimbraId"})

	if err != nil {
		return "", err
	}

	if len(response.Content.Cos) == 0 || response.Content.Cos[0].ID == "" {
		return "", fmt.Errorf("no such cos: %s", key)
	}

	return response.Content.Cos[0].ID, nil
}

func validateCosName(name string) error {
	if name == "" || strings.ContainsFunc(name, unicode.IsSpace) {
		return fmt.Errorf("invalid cos name %q", name)
	}

	return nil
}

func validateDomainName(name string) error {
	if name == "" || strings.Contains(name, "@") || strings.ContainsFunc(name, unicode.IsSpace) {
		return fmt.Errorf("invalid domain name %q", name)
	}

	return nil
}

func validateAddress(address string) error {
	local, domain, found := strings.Cut(address, "@")

	if !found || local == "" || domain == "" || strings.Contains(domain, "@") || strings.ContainsFunc(address, unicode.IsSpace) {
		return fmt.Errorf("invalid address %q: expected name@domain", address)
	}

	return nil
}

func validateAddresses(addresses []string) error {
	if len(addresses) == 0 {
		return fmt.Errorf("no address given")
	}

	for _, address := range addresses {
		if err := validateAddress(address); err != nil {
			return err
		}
	}

	return nil
}

func buildAttrs(args []string, keepModifiers bool) (zimbraCommon.AttrsNode, error) {
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("attribute %q has no value", args[len(args)-1])
	}

	attrs := make(zimbraCommon.AttrsNode, 0, len(args)/2)

	for i := 0; i < len(args); i += 2 {
		name := args[i]
		base := name

		if len(name) > 0 && (name[0] == '+' || (keepModifiers && name[0] == '-')) {
			base = name[1:]

			if !keepModifiers {
				name = base
			}
		}

		if base == "" {
			return nil, fmt.Errorf("empty attribute name")
		}

		attrs = append(attrs, zimbraCommon.AttrNode{
			Name:  name,
			Value: args[i+1],
		})
	}

	return attrs, nil
}
