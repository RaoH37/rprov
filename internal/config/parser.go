package config

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type LocalConfig struct {
	XMLName xml.Name `xml:"localconfig"`
	Keys    []Key    `xml:"key"`
}

type Key struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

func ParseLocalConfig(filePath string) (map[string]string, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var config LocalConfig
	err = xml.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	configMap := make(map[string]string)
	for _, k := range config.Keys {
		configMap[k.Name] = k.Value
	}

	return configMap, nil
}
