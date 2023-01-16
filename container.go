package kata_hook_api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	spec "github.com/opencontainers/runtime-spec/specs-go"
)

type containerConfig struct {
	version     string
	Pid         int
	Rootfs      string
	Envs        map[string]string
	Annotations map[string]string
}

func loadContainerSpec() (*containerConfig, error) {
	var state spec.State
	if err := json.NewDecoder(os.Stdin).Decode(&state); err != nil {
		log.Println("could not decode container state")
		return nil, err
	}

	bundle := state.Bundle
	configFilePath := path.Join(bundle, "config.json")
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Println("could not open config file")
		return nil, err
	}
	defer file.Close()

	var spec *spec.Spec
	if err = json.NewDecoder(file).Decode(&spec); err != nil {
		log.Println("could not decode oci spec")
		return nil, err
	}

	if spec.Version == "" || spec.Process == nil || spec.Root == nil {
		return nil, fmt.Errorf("oci spec parameter error")
	}

	env, err := getEnv(spec.Process.Env)
	if err != nil {
		log.Println("could not parse container env")
		return nil, err
	}

	annotations := make(map[string]string)
	for k, v := range spec.Annotations {
		annotations[k] = v
	}

	return &containerConfig{
		version:     spec.Version,
		Pid:         state.Pid,
		Rootfs:      spec.Root.Path,
		Envs:        env,
		Annotations: annotations,
	}, nil
}

func getEnv(env []string) (map[string]string, error) {
	envs := make(map[string]string)
	log.Printf("env list: %v", env)
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid environment variable: %v", e)
		}
		envs[parts[0]] = parts[1]
	}
	return envs, nil
}
