package goutil

import (
	"errors"
	// "log"
	"testing"
)

func TestGetStringsFromJson(t *testing.T) {

	// items, err := GetStringsFromJson("testdata/test.json")
	_, err := GetStringsFromJson("testdata/test.json")
	if err != nil {
		t.Error(err)
	}
	// log.Println(items)

	_, err = GetStringsFromJson("testdata/xxx.json")
	if err == nil {
		t.Error(err)
	}
	// log.Println(err)
}

func TestIsExist(t *testing.T) {

	b := IsExist("testdata/test.json")
	if !b {
		t.Error(errors.New("file is exist,test is fail"))
	}

	b = IsExist("testdata/xxx.json")
	if b {
		t.Error(errors.New("file is not exist,test is fail"))
	}
}

func TestGetImagesLinks(t *testing.T) {
	// links, err := GetImagesLinks("http://www.163.com")
	_, err := GetImagesLinks("http://www.163.com")
	if err != nil {
		t.Error(err)
	}
	// log.Println(links)
}

func TestGetHyperlinks(t *testing.T) {

	// links, err := GetHyperlinks("http://www.163.com")
	_, err := GetHyperlinks("http://www.163.com")
	if err != nil {
		t.Error(err)
	}
	// log.Println(links)
}
