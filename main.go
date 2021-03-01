package main

import (
	"fmt"
	"log"
	"workshop-golang-psm/food-app-service/api"
	"workshop-golang-psm/food-app-service/pkg/config"
	"workshop-golang-psm/food-app-service/pkg/database"

	"github.com/spf13/viper"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))

	app := api.SetupRouter(db)
	app.Run(port)
}
