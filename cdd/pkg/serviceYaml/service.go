package serviceYaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetServiceYAML(filepath string) (ServiceYAML, error) {
	var result ServiceYAML
	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	err = yaml.Unmarshal(filedata, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
