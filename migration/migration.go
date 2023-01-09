package main

import (
	"fmt"
	"github.com/cnugroho11/url_shortener/initializers"
	"github.com/cnugroho11/url_shortener/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env")
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Url{})

	fmt.Println("Migration complete")
}
