package csvConvert

import (
	"testing"
)

// TestOpenCSV check to open file and content
func TestOpenCSV(t *testing.T) {

	// test.csv used for this test
	file := "./test.csv"
	data, err := OpenCSV(file)
	if err != nil {
		t.Fatal(err)
	}

	// check content data
	if len(data) < 1 {
		t.Fatal("no data inside test.csv")
	}
}
