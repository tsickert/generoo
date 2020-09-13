package registry

import (
	"fmt"
	"github.com/army-of-one/generoo/pkg/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	DefaultRegistry       = ".generoo/registry"
	RegistryLinkManifest  = ".generoo/registry/links.yml"
	GenerooRegistryEnvVar = "GENEROO_REGISTRY"
)

func Register(target string) (string, error) {
	isGeneroo, err := utils.HasGenerooDir(target)
	if err != nil || !isGeneroo {
		log.Fatalf("could not find a .generoo directory: %s", err)
		return "", err
	}

	registryDir, exists := os.LookupEnv(GenerooRegistryEnvVar)
	if !exists {
		registryDir = filepath.Join(utils.PathFromHome(DefaultRegistry), filepath.Base(target))
	}

	err = utils.Copy(target, registryDir)
	if err != nil {
		log.Fatalf("failed to register %s", target)
	}

	return filepath.Base(target), nil
}

func RegisterLocal() error {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Fatal("failed to get the current working directory")
	}

	return Register(workingDir)
}

// Link leverages OS level symlinks to point to an existing directory
// Note: this is an experimental feature
func Link(target string) error {
	isGeneroo, err := utils.HasGenerooDir(target)
	if err != nil || !isGeneroo {
		log.Fatalf("could not find a .generoo directory: %s", err)
		return err
	}

	alias := filepath.Base(target)

	if alias == "." {
		err = LinkLocal()
	}

	newLink, exists := os.LookupEnv(GenerooRegistryEnvVar)
	if !exists {
		newLink = filepath.Join(utils.PathFromHome(DefaultRegistry), RegistryLinkManifest, filepath.Base(target))
	}

	err = os.Link(target, newLink)

	return nil
}

func LinkLocal() error {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Fatal("failed to get the current working directory")
	}

	return Link(workingDir)
}

func List() error {
	var err error
	var fds []os.FileInfo

	generooDir := utils.PathFromHome(DefaultRegistry)

	if fds, err = ioutil.ReadDir(generooDir); err != nil {
		return err
	}

	for _, fd := range fds {
		if fd.IsDir() {
			fmt.Print(fd.Name())
		}
	}

	return nil
}
