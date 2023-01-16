package kata_hook_api

import (
	"flag"
	"fmt"
)

//type InitFunc func(context.Context, interface{}) error

type Access interface {
	Execute() error
}

type config struct {
	container *containerConfig
}

func NewAccess(h HookHandlers) (Access, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return nil, fmt.Errorf("incorrect command args")
	}

	containerConf, err := loadContainerSpec()
	if err != nil {
		return nil, err
	}

	return &controller{
		cfg: &config{
			container: containerConf,
		},
		option:    args[0],
		handlers:  h,
	}, nil
}

func MetaEnv(obj interface{}, key string) (string, error) {
	conf, ok := obj.(*config)
	if ok && conf.container != nil {
		if val, ok := conf.container.Envs[key]; ok {
			return val, nil
		} else {
			return "", fmt.Errorf("env not found: %s", key)
		}
	}

	return "", fmt.Errorf("object has no meta")
}

func MetaAnnotation(obj interface{}, key string) (string, error) {
	conf, ok := obj.(*config)
	if ok && conf.container != nil {
		if val, ok := conf.container.Annotations[key]; ok {
			return val, nil
		} else {
			return "", fmt.Errorf("annotation not found: %s", key)
		}
	}

	return "", fmt.Errorf("object has no meta")
}