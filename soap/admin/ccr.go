package admin

import (
	"fmt"
	"strings"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

func CreateCalendarResource(conf map[string]string, name string, password string, attributes []string) error {
	if err := validateAddress(name); err != nil {
		return err
	}

	attrs, err := buildAttrs(attributes, false)

	if err != nil {
		return err
	}

	if err := validateCalendarResourceAttrs(attrs); err != nil {
		return err
	}

	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	response, err := zcs.CreateCalendarResourceRequest(name, password, attrs)

	if err != nil {
		return err
	}

	if len(response.Content.Calresource) == 0 {
		return fmt.Errorf("server returned no calendar resource after creating %s", name)
	}

	fmt.Printf("%v\n", response.Content.Calresource[0].ID)

	return nil
}

func validateCalendarResourceAttrs(attrs zimbraCommon.AttrsNode) error {
	hasDisplayName := false
	hasType := false

	for _, attr := range attrs {
		switch {
		case strings.EqualFold(attr.Name, "displayName"):
			hasDisplayName = attr.Value != ""
		case strings.EqualFold(attr.Name, "zimbraCalResType"):
			if attr.Value != "Location" && attr.Value != "Equipment" {
				return fmt.Errorf("invalid zimbraCalResType %q: expected Location or Equipment", attr.Value)
			}

			hasType = true
		}
	}

	if !hasDisplayName {
		return fmt.Errorf("displayName is required")
	}

	if !hasType {
		return fmt.Errorf("zimbraCalResType is required (Location or Equipment)")
	}

	return nil
}
