package admin

import (
	"fmt"
	"strings"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

var aliasDomainAttrs = []string{
	"zimbraDomainType",
	"zimbraDomainAliasTargetId",
	"zimbraMailCatchAllAddress",
	"zimbraMailCatchAllForwardingAddress",
}

func CreateAliasDomain(conf map[string]string, aliasName string, targetKey string, attributes []string) error {
	if err := validateDomainName(aliasName); err != nil {
		return err
	}

	attrs, err := buildAttrs(attributes, false)

	if err != nil {
		return err
	}

	for _, attr := range attrs {
		for _, reserved := range aliasDomainAttrs {
			if strings.EqualFold(attr.Name, reserved) {
				return fmt.Errorf("attribute %s is set by cad and cannot be given", reserved)
			}
		}
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, name := splitIdOrName(targetKey)

	targets, err := zcs.GetDomainRequest(id, name, zimbraCommon.StringList{"zimbraId"}, false)

	if err != nil {
		return err
	}

	if len(targets.Content.Domains) == 0 || targets.Content.Domains[0].ID == "" || targets.Content.Domains[0].Name == "" {
		return fmt.Errorf("no such domain: %s", targetKey)
	}

	target := targets.Content.Domains[0]

	attrs = append(attrs,
		zimbraCommon.AttrNode{Name: "zimbraDomainType", Value: "alias"},
		zimbraCommon.AttrNode{Name: "zimbraDomainAliasTargetId", Value: target.ID},
		zimbraCommon.AttrNode{Name: "zimbraMailCatchAllAddress", Value: "@" + aliasName},
		zimbraCommon.AttrNode{Name: "zimbraMailCatchAllForwardingAddress", Value: "@" + target.Name},
	)

	response, err := zcs.CreateDomainRequest(aliasName, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Domain) == 0 {
		return fmt.Errorf("server returned no domain after creating %s", aliasName)
	}

	fmt.Printf("%v\n", response.Content.Domain[0].ID)

	return nil
}
