package clients

import (
	"encoding/json"
	"fmt"
	"scraper/config"
	"scraper/services"
	"time"
)

type BondsClient struct {
	BondsDataSourceURL string
	BondsViewURL       string
}

type Offer struct {
	InterestRate float64 `json:"interestRate"`
	Period       int     `json:"period"`
}

type OffersResponse []*Offer

func NewBondsClient() *BondsClient {
	config := config.GetConfig()

	return &BondsClient{
		BondsDataSourceURL: config.BondsDataSourceURL,
		BondsViewURL:       config.BondsViewURL,
	}
}

func (c *BondsClient) getBondsOffers() (*OffersResponse, error) {
	offersResponse := &OffersResponse{}
	if err := services.GetRequest(c.BondsDataSourceURL, offersResponse); err != nil {
		fmt.Println("[getBondsOffers] Failed to get data")

		return nil, err
	}

	responseJsonString, err := json.Marshal(offersResponse)
	if err != nil {
		fmt.Println("[getBondsOffers] Failed to marshal response", err)
	} else {
		fmt.Println("Offers data", string(responseJsonString))
	}

	return offersResponse, nil
}

func (c *BondsClient) prepareAndSendEmail(interestRate float32) error {
	timeNow := time.Now()
	emailClient := services.NewEmailClient()
	subject := "Krājobligāciju procentu likmes"
	message := "<h3> 12 mēnēšu krājobligāciju procentu likme šodien ir sasniegusi vēlamo līmeni! </h3> <br/>" +
		"<p> Uz " + timeNow.Format("02.01.2006 15:04") + " likme ir: <strong>" + fmt.Sprintf("%.2f", interestRate) + "</strong></p> <br/>" +
		"<p> <a href='" + c.BondsViewURL + "'>Atvērt piedāvājumu</a></p>"

	if err := emailClient.SendEmail(subject, message); err != nil {
		fmt.Println("[ProcessSavingBondsOffers] Failed to send email", err)

		return err
	}

	return nil
}

func (c *BondsClient) ProcessSavingBondsOffers() error {
	bondOffers, err := c.getBondsOffers()
	if err != nil {
		return err
	}

	for _, offer := range *bondOffers {
		if offer.Period == 12 {
			if err = c.prepareAndSendEmail(float32(offer.InterestRate)); err != nil {
				return err
			}
		}
	}

	return nil
}
