package main

import "github.com/cripplemymind9/go-market/internal/app"

const configPath = "config/config.yaml"

// @title           Go-market
// @version         1.0
// @description     This service manages product purchases and provides endpoints to interact with product and purchase data.

// @contact.name   Egor K.
// @contact.email  ololoevlan@gmail.com

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	app.Run(configPath)
}
