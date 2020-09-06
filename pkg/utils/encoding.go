package utils

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
)

// ReadAmbiguous attempts to read files in either json or yaml format. If a file is determined to be yaml, it is
// converted to json before being returned.
func ReadAmbiguousAsJson(file string) ([]byte, error) {
	if content, err := ioutil.ReadFile(file); err == nil {
		var holder map[string]interface{}
		if err = json.Unmarshal(content, &holder); err != nil {
			if err = yaml.Unmarshal(content, &holder); err != nil {
				log.Printf("unable to read %s, please provide in valid json or yaml format", file)
				return nil, err
			} else {
				if content, err = yaml.YAMLToJSON(content); err != nil {
					return nil, err
				}
			}
		}
		return content, nil
	} else {
		return nil, err
	}
}
