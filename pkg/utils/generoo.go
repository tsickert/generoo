package utils

import (
	"io/ioutil"
	"log"
)

func HasGenerooDir(dir string) (bool, error) {
	if files, err := ioutil.ReadDir(dir); err != nil {
		log.Fatal("failed to read the current working directory")
		return false, err
	} else {
		for _, f := range files {
			if f.Name() == ".generoo" {
				log.Printf("found generoo directory, adding to registry")
				return true, nil
			}
		}
		log.Printf("could not find generoo directory, failed to add to registry")
		return false, nil
	}
}

func GetGenerooMetadata(config string) {

}