package main

import (
	"log"

	config "github.com/profile/service/pkg/config"
	di "github.com/profile/service/pkg/di"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatalln("Could not load the config file :", cfgErr)
		return
	}
	server, err := di.InitApi(cfg)
	if err != nil {
		log.Fatalln("Error in initializing the API", err)
	}
	server.Start()

}
