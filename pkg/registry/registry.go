package registry

import (
	"github.com/army-of-one/generoo/pkg/utils"
	"log"
	"os"
	"path"
)

const (
	DefaultRegistry       = ".generoo/registry"
	GenerooRegistryEnvVar = "GENEROO_REGISTRY"
)

func Register(target string) error {
	isGeneroo, err := utils.HasGenerooDir(target)
	if err != nil || !isGeneroo {
		log.Fatalf("could not find a .generoo directory: %s", err)
		return err
	}

	registryDir, exists := os.LookupEnv(GenerooRegistryEnvVar)
	if !exists {
		registryDir = path.Join(utils.PathFromHome(DefaultRegistry), path.Base(target))
	}

	err = utils.Copy(target, registryDir)
	if err != nil {
		log.Fatalf("failed to register %s", target)
	}

	return nil
}

func RegisterLocal() error {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Fatal("failed to get the current working directory")
	}

	return Register(workingDir)
}
