package clients

import (
	"encoding/json"
	"fmt"
	"scraper/services"
)

type BondsClient struct {
	DataSourceURL string
}

type Offer struct {
	InterestRate float64 `json:"interestRate"`
	Period       int     `json:"period"`
}

type OffersResponse []*Offer

func NewBondsClient() *BondsClient {
	return &BondsClient{
		DataSourceURL: "https://www.krajobligacijas.lv/api/offers/today",
	}
}

func (c *BondsClient) GetBondsOffers() error {
	offersResponse := &OffersResponse{}
	if err := services.GetRequest(c.DataSourceURL, offersResponse); err != nil {
		fmt.Println("Failed to get data")
	}

	responseJsonString, err := json.Marshal(offersResponse)
	if err != nil {
		fmt.Println("Failed to marshal response", err)

		return err
	}

	fmt.Println(string(responseJsonString))

	return nil
}
