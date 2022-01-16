package main

import (
	"flag"
	apiserver "github.com/BEEugene/Go433MhzGrabber/internal/app"
	"github.com/BurntSushi/toml"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to configs")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
