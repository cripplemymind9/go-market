package main

import "github.com/cripplemymind9/go-market/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}