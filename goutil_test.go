package goutil

import (
	// "log"
	"testing"
)

func TestGetStringsFromJson(t *testing.T) {

	_, err := GetStringsFromJson("testdata/test.json")
	if err != nil {
		t.Error(err)
	}
	// log.Println(items)
}
