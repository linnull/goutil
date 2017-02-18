package goutil

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	Data []string
}

func GetStringsFromJson(jsonName string) ([]string, error) {
	b, err := ioutil.ReadFile(jsonName)
	if err != nil {
		return nil, err
	}
	items := &Data{[]string{}}
	err = json.Unmarshal(b, items)
	if err != nil {
		return nil, err
	}
	return items.Data, nil

}
