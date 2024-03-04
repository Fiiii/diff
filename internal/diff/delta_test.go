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
		t.Logf("\tTest %d:\tWhen calculating chunks hashes.", testID)
		{
			originalData := []byte("hello")
			expHashes := []uint64{0}
			chunks := divideIntoChunks(testChunkSize, originalData)
			hashes := calculateRollingHashes(chunks)

			if reflect.DeepEqual(hashes, expHashes) {
				t.Logf("\t%s\tTest %d:\tShould be able to return proper divided chunks.", Success, testID)
			} else {
				t.Logf("got: %v", hashes)
				t.Logf("exp: %v", expHashes)
				t.Fatalf("\t%s\tTest %d:\tShould be able to return proper divided chunks.", Failed, testID)
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
			testID := i + 1
			t.Logf("\tTest %d:\tWhen handling delta calculation for original: `%s` and updated: `%s`.",
				testID, tc.originalData, tc.updatedData)
			{

			}
		}
	}
}
