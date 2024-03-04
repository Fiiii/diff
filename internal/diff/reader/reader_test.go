package reader

import (
	"os"
	"testing"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestReader(t *testing.T) {
	t.Log("Given the need to work with file reader")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen reading file.", testID)
		{
			tmp, err := os.CreateTemp("", "tmp")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmp.Name())

			testData := []byte("Hello, World!")
			if _, err := tmp.Write(testData); err != nil {
				t.Fatal(err)
			}

			if err := tmp.Close(); err != nil {
				t.Fatal(err)
			}

			data, err := ReadFromPath(tmp.Name())
			if err != nil {
				t.Errorf("ReadFromPath(%s) returned an error: %v", tmp.Name(), err)
			}

			if string(data) != string(testData) {
				t.Logf("got: %v", string(data))
				t.Logf("exp: %v", string(testData))
				t.Fatalf("\t%s\tTest %d:\tShould be able to return text from file.", Failed, testID)
			} else {
				t.Logf("\t%s\tTest %d:\tShould be able to read properly from file.", Success, testID)
			}
		}
	}
}
