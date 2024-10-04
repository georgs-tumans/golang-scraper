package main

import (
	"log"
	"os"
	"scraper/clients"
)

func main() {
	logFile := setupLogger()
	defer logFile.Close()

	log.Println("Running scraper")
	bondsClient := clients.NewBondsClient()
	if err := bondsClient.ProcessSavingBondsOffers(); err != nil {
		log.Println("Failed to process bonds data")
	}

}

func setupLogger() *os.File {
	file, err := os.OpenFile("golang_web_scraper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[setupLogger] Failed to open log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return file
}
