// Package reader provides a way to read the file from the given path.
package reader

import "os"

// ReadFromPath reads the file from the given path
func ReadFromPath(path string) ([]byte, error) {
	originalData, err := os.ReadFile("cmd/" + path)
	if err != nil {
		return nil, err
	}

	return originalData, nil
}
