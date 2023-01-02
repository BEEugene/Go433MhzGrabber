package main

import (
	"github.com/BEEugene/Go433MhzGrabber/docs" // docs is generated by Swag CLI, you have to import it.
	"flag"
	"log"
	"os"

	apiserver "github.com/BEEugene/Go433MhzGrabber/internal/app"
	"github.com/BurntSushi/toml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./config/apiserver.toml", "path to configs")
}

// @title Orders API
// @contact.name API Support
func main() {
	flag.Parse()
	print(os.Getwd())

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	
	docs.SwaggerInfo.Host = "localhost" + config.BindAddr
	docs.SwaggerInfo.BasePath = ""
	if err != nil {
		log.Fatal(err)
	}
	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
