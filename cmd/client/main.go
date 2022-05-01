package main

import "vladazn/wow/internal/client"

const configPath = "config/config.yml"

func main() {
	err := client.Run(configPath)
	if err != nil {
		panic(err)
	}
}
