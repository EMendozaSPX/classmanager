package Env

import (
	"os"
)

func exists(path string) (bool, error) {
	// Check if directory or file exists returns an error
	_, err := os.Stat(path)

	// If no errors occurred file exists so return true
	if err == nil {
		return true, nil
	}

	// If file not found error occurred return false
	if os.IsNotExist(err) {
		return false, nil
	}

	// Return false with error message in all other cases
	return false, err
}



func import_variables() {

}
