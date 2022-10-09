package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

func TestConfig(t *testing.T) {
	config := &Conf{}
	_, err := toml.DecodeFile("conf.toml", config)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%+v", config))
}
