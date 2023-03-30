package main

import "net/http"

type SearchQuery struct {
	Kind       string `json:"kind"`
	TotalItems int64  `json:"totalItems"`
	Items      Items  `json:"items"`
}

type Items struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
}

type SaleInfo struct {
	Country     string      `json:"country"`
	Saleability string      `json:"saleability"`
	IsEbook     bool        `json:"isEbook"`
	ListPrice   ListPrice   `json:"listPrice"`
	RetailPrice RetailPrice `json:"retailPrice"`
	BuyLink     string      `json:"buyLink"`
}

type ListPrice struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type RetailPrice struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type Offers []struct {
	FinskyOfferType  int64            `json:"finskyOfferType"`
	OfferListPrice   OfferListPrice   `json:"listPrice"`
	OfferRetailPrice OfferRetailPrice `json:"listPrice"`
	Gifatable        bool             `json:"giftable"`
}

type OfferListPrice struct {
	AmountInMicros int64  `json:"amountInMicros"`
	CurrencyCode   string `json:"currencyCode"`
}
type OfferRetailPrice struct {
	AmountInMicros int64  `json:"amountInMicros"`
	CurrencyCode   string `json:"currencyCode"`
}
type AccessInfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"veiwability"`
	Embeddable             bool   `json:"embeddable"`
	PublicDomain           bool   `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	Epub                   Epub   `json:"epub"`
	PDF                    PDF    `json:"pdf"`
	WebReaderLink          string `json:"webReaderLink"`
	AccessViewStatus       string `json:"accessViewStatus"`
	QuoteSharingAllowed    bool   `json:"quoteSharingAllowed"`
}
type Epub struct {
	IsAvailable bool `json:"isAvailable"`
}
type PDF struct {
	IsAvailable bool `json:"isAvailable"`
}

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

func (app *Config) Search(w http.ResponseWriter, r *http.Request) {

}
