// Package diff provides a way to generate a delta between two byte slices.
package diff

const (
	// primeRK is the prime base used in Rabin-Karp algorithm.
	primeRK = 16777619
)

type Delta struct {
	ChunksToReuse []Chunk
	Changes       []Change
	Removals      []Chunk
	Additions     []Chunk
}

type Change struct {
	Position int
	OldData  []byte
	NewData  []byte
}

type Chunk struct {
	Position int
	Data     []byte
}

// GenerateDelta generates a delta between the original and updated byte slices.
func GenerateDelta() (Delta, error) {
	return Delta{}, nil
}

// divideIntoChunks divides the data into chunks of size chunkSize.
func divideIntoChunks(chunkSize int, data []byte) []Chunk {
	var chunks []Chunk

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, Chunk{Position: i, Data: data[i:end]})
	}
	return chunks
}

// calculateRollingHashes calculates the rolling hashes for each chunk.
func calculateRollingHashes(chunks []Chunk) []uint64 {
	var hashes []uint64
	for i, _ := range chunks {
		var rollingHash uint64
		for j, _ := range chunks[i].Data {
			rollingHash = rollingHash*primeRK + uint64(chunks[i].Data[j])
		}
		hashes = append(hashes, rollingHash)
	}
	return hashes
}

// compareHashes compares the hashes of the original and updated chunks to generate a delta.
func compareHashes() Delta {
	return Delta{}
}
