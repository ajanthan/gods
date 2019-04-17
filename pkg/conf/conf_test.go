package conf

import (
	"fmt"
	"testing"
)

func TestLoadGoDs(t *testing.T) {
	gods, err := LoadGoDs("../../spec/sample/sample-1.yaml")
	if err != nil {
		t.Error("Error in loading conf ", err)
	}
	if gods.Name != "CartAPI" {
		t.Error("Conf name doesn't match")
	}
	fmt.Printf("Conf:%+v ", gods)
}
