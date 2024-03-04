package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fiiii/diff/internal/diff"
	"github.com/fiiii/diff/internal/diff/reader"
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

	for _, chunk := range delta.ChunksToReuse {
		fmt.Printf("Chunk at position %d is reusable\n", chunk.Position)
	}

	for _, change := range delta.Changes {
		fmt.Printf("Change at position %d:\n", change.Position)
		fmt.Printf("  Old Data: %s\n", string(change.OldData))
		fmt.Printf("  New Data: %s\n", string(change.NewData))
	}

	for _, removal := range delta.Removals {
		fmt.Printf("Chunk at position %d removed\n", removal.Position)
	}

	for _, addition := range delta.Additions {
		fmt.Printf("Chunk at position %d added\n", addition.Position)
	}
}
