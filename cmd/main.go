package main

import "fmt"

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
}
