package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/ajanthan/gods/pkg/model"
)

func LoadGoDs(fileLocation string) (model.Gods, error) {
	var gods model.Gods
	btyes, fileErr := ioutil.ReadFile(fileLocation)
	if fileErr != nil {
		return gods, fileErr
	}
	unMarshalErr := yaml.Unmarshal(btyes, &gods)
	if unMarshalErr != nil {
		return gods, unMarshalErr
	}
	return gods, nil
}
