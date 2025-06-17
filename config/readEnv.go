package config

import (
	"errors"
	"os"
)

// ReadEnv reads an environment variable.
// If the variable is not set, it returns the provided defaultValue if available,
// otherwise, it returns an error.
// The defaultValue is optional.
func ReadEnv(str string, defaultValue ...string) (string, error) {
	val, exists :=  os.LookupEnv(str)

	if !exists {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}

		err := errors.New("variable does not exist")
		return "", err
	}
	return val, nil
}