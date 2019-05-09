package config

import (
	"fmt"
	"testing"
)

func Test_populateConfig(t *testing.T) {
	if err := populateConfig("./../config.json"); err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(Config)
}
