package main

import (
	"fmt"
	"scraper/clients"
)

func main() {
	bondsClient := clients.NewBondsClient()
	if err := bondsClient.GetBondsOffers(); err != nil {
		fmt.Println("Failed to process bonds data")
	}

}
