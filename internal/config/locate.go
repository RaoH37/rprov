package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var localConfigPaths = []string{
	"/opt/zimbra/conf/localconfig.xml",
	"/opt/zextras/conf/localconfig.xml",
}

func FindLocalConfig() (string, error) {
	for _, p := range localConfigPaths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
	}

	return "", fmt.Errorf("localconfig.xml not found (searched: %s)", strings.Join(localConfigPaths, ", "))
}
