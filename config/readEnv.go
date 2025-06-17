package config

import (
	"errors"
	"os"
)

func ReadEnv(str string) (string, error) {
	val, exists :=  os.LookupEnv(str)

	if !exists {
		err := errors.New("variable does not exist")
		return val, err
	}
	return val, nil
}