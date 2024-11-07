package utils

import (
	"fmt"
	"os"
)

/*
given a env variable name, returns the values for that key

params:

	key string

returns:

	string, error
*/
func GetEnv(key string) (string, error) {
	val := os.Getenv(key)

	if len(val) == 0 {
		return "", fmt.Errorf("could not find env variable with key %s", key)
	}

	return val, nil
}
