package diff

import (
	"reflect"
	"testing"
)

const (
	Success       = "\u2713"
	Failed        = "\u2717"
	testChunkSize = 3
)

func TestDelta(t *testing.T) {
	t.Log("Given the need to work with chunks division")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen dividing into chunks.", testID)
		{
			originalData := []byte("hello")
			// Assuming 3 is the chunk size
			expChunks := []Chunk{{0, []byte("hel")}, {3, []byte("lo")}}
			chunks := divideIntoChunks(testChunkSize, originalData)

			if reflect.DeepEqual(chunks, expChunks) {
				t.Logf("\t%s\tTest %d:\tShould be able to return proper divided chunks.", Success, testID)
			} else {
				t.Logf("got: %v", chunks)
				t.Logf("exp: %v", expChunks)
				t.Fatalf("\t%s\tTest %d:\tShould be able to return proper divided chunks.", Failed, testID)
			}
		}

		testID = 1
		t.Logf("\tTest %d:\tWhen calculating rolling hashes.", testID)
		{
			originalData := []byte("hello")
			// Based on "good" hashing prime number
			expHashes := []uint64{29274805622692371, 1811982963}
			chunks := divideIntoChunks(testChunkSize, originalData)
			hashes := calculateRollingHashes(chunks)

			if reflect.DeepEqual(hashes, expHashes) {
				t.Logf("\t%s\tTest %d:\tShould be able to return calculated hashed chunks.", Success, testID)
			} else {
				t.Logf("got: %v", hashes)
				t.Logf("exp: %v", expHashes)
				t.Fatalf("\t%s\tTest %d:\tShould be able to return calculated hashed chunks.", Failed, testID)
			}
		}
	}

	t.Log("Given the need to work with delta calculations")
	{
		var deltaTestsCases = []struct {
			originalData []byte
			updatedData  []byte
			delta        Delta
		}{
			{originalData: []byte("abc"), updatedData: []byte("abc"), delta: Delta{
				ChunksToReuse: []Chunk{{0, []byte("abc")}}}},
			{originalData: []byte("abc"), updatedData: []byte("abcd"), delta: Delta{
				ChunksToReuse: []Chunk{{0, []byte("abc")}},
				Additions:     []Chunk{{3, []byte("d")}}}},
			{originalData: []byte("abcd"), updatedData: []byte("abc"), delta: Delta{
				ChunksToReuse: []Chunk{{0, []byte("abc")}},
				Removals:      []Chunk{{3, []byte("d")}}}},
			{originalData: []byte("abc"), updatedData: []byte("def"), delta: Delta{
				Changes: []Change{{0, []byte("abc"), []byte("def")}}}},
		}
		for i, tc := range deltaTestsCases {
			testID := i + 2
			t.Logf("\tTest %d:\tWhen handling delta calculation for original: `%s` and updated: `%s`.",
				testID, tc.originalData, tc.updatedData)
			{
				t.Fatalf("\t%s\tTest %d:\tShould be able to return proper calculated delta.", Failed, testID)
			}
		}
	}
}
