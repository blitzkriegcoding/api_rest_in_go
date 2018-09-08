package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	var gopath = os.Getenv("GOPATH")
	if _, err := toml.DecodeFile(gopath+"/src/github.com/blitzkriegcoding/api_go/config.toml", &c); err != nil {
		log.Fatal(err)
	}
	log.Println(gopath + "/src/github.com/blitzkriegcoding/api_go/config.toml")
}
