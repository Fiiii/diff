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
func GenerateDelta(chunkSize int, original, updated []byte) (Delta, error) {
	originalChunks := divideIntoChunks(chunkSize, original)
	updatedChunks := divideIntoChunks(chunkSize, updated)

	originalHashes := calculateRollingHashes(originalChunks)
	updatedHashes := calculateRollingHashes(updatedChunks)

	return compareHashes(originalChunks, updatedChunks, originalHashes, updatedHashes), nil
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
func compareHashes(originalChunks, updatedChunks []Chunk, originalHashes, updatedHashes []uint64) Delta {
	var delta Delta

	minLength := len(originalHashes)
	if len(updatedHashes) < minLength {
		minLength = len(updatedHashes)
	}

	for i := 0; i < minLength; i++ {
		if originalHashes[i] == updatedHashes[i] {
			delta.ChunksToReuse = append(delta.ChunksToReuse, originalChunks[i])
		} else {
			delta.Changes = append(delta.Changes, Change{
				Position: originalChunks[i].Position,
				OldData:  originalChunks[i].Data,
				NewData:  updatedChunks[i].Data,
			})
		}
	}

	for i := minLength; i < len(originalHashes); i++ {
		delta.Removals = append(delta.Removals, originalChunks[i])
	}

	for i := minLength; i < len(updatedHashes); i++ {
		delta.Additions = append(delta.Additions, updatedChunks[i])
	}

	return delta
}
