package utils

import (
	"os"
)

// GetEnvVar retrieves a property from the Viper object or returns a default value if it is not set.
//
// Params:
//
// lookUpVar (string): The name of the environment variable to retrieve.
// defaultVal (string): The default value to return if the environment variable is not set.
//
// Returns:
//
// string: The value of the environment variable or the default value if it is not set.
// bool: True if the value was set.
func GetEnvVar(lookUpVar, defaultVal string) (string, bool) {
	val, ok := os.LookupEnv(lookUpVar)
	if !ok {
		return defaultVal, false
	} else {
		return val, true
	}
}
