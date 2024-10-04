package main

import (
	"fmt"
	"scraper/clients"
)

func main() {
	bondsClient := clients.NewBondsClient()
	if err := bondsClient.ProcessSavingBondsOffers(); err != nil {
		fmt.Println("Failed to process bonds data")
	}

}
