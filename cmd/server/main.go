package main

import "vladazn/wow/internal/server"

const configPath = "config/config.yml"

func main() {
	err := server.Run(configPath)
	if err != nil {
		panic(err)
	}
}
