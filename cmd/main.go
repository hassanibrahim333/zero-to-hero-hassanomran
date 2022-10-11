package main

import (
	"fmt"
	"secondchallange/Internal/Repository"
	"secondchallange/Internal/adapter/Api"
	"secondchallange/Internal/adapter/dbConnections"
	"secondchallange/Internal/service"
	"secondchallange/config"
)

func main() {
	var configuration config.Configurations
	config, err := config.SetUpViper(configuration)
	if err != nil {
		fmt.Println("Failed to read Config file")
	}
	con, err := dbConnections.ConnectToDatabase(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	tranRepo := Repository.NewDefaultRepository(con)
	tranSer := service.NewDefaultService(tranRepo)
	Api.Request(tranSer, config)
}
