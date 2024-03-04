package main

import (
	"fmt"
	"github.com/fiiii/diff/internal/reader"
	"os"
	"strconv"

	"github.com/fiiii/diff/internal/diff"
)

const (
	defaultChunkSize = 3
	defaultFile1     = "cmd/original"
	defaultFile2     = "cmd/updated"
)

// Hashing function gets the data as a parameter. Separate possible filesystem operations.
// Chunk size can be fixed or dynamic, but must be split to at least two chunks on any sufficiently sized data.
// Should be able to recognize changes between chunks. Only the exact differing locations should be added to the delta.
// TODO:
// 1. Read []bytes from files/input streams
// 2. Divide into pieces/chunks
// 3. Calculate rolling hashes per each chunk
// 4. Compare hashes - calculate delta
func main() {
	fmt.Println("Hello, rolling hash!")
	var chunkSize int
	var file1, file2 string
	if len(os.Args) == 4 {
		// Parse command-line arguments
		cs, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid chunk size, using default value = 4", err)
			chunkSize = defaultChunkSize
		} else {
			chunkSize = cs
		}

		if err != nil {
			fmt.Println("Error getting working directory:", err)
			return
		}

		file1 = os.Args[2]
		file2 = os.Args[3]
	} else {
		// Use default values
		chunkSize = defaultChunkSize
		file1 = defaultFile1
		file2 = defaultFile2
	}

	original, err := reader.ReadFromPath(file1)
	if err != nil {
		fmt.Println("Reader error:", err)
		return
	}

	updated, err := reader.ReadFromPath(file2)
	if err != nil {
		fmt.Println("Reader error:", err)
		return
	}

	delta, err := diff.GenerateDelta(chunkSize, original, updated)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	delta.Print()
}
