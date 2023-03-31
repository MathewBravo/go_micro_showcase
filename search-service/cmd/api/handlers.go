package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SearchQuery struct {
	Kind       string `json:"kind"`
	TotalItems int64  `json:"totalItems"`
	Items      Items  `json:"items"`
}

type Items []struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

type VolumeInfo struct {
	Title               string              `json:"title"`
	Subtitle            string              `json:"subtitle"`
	Authors             []string            `json:"authors"`
	Publisher           string              `json:"publisher"`
	PublishedDate       string              `json:"publishedDate"`
	Description         string              `json:"description"`
	IndustryIdentifiers IndustryIdentifiers `json:"industryIdentifiers"`
	ReadingModes        ReadingModes        `json:"readingModes"`
	PageCount           int64               `json:"pageCount"`
	PrintType           string              `json:"printType"`
	Categories          []string            `json:"categories"`
	AverageRating       int64               `json:"averageRating"`
	RatingsCount        int64               `json:"ratingsCount"`
	MaturityRating      string              `json:"maturityRating"`
	AllowAnonLogging    bool                `json:"allowAnonLogging"`
	ContentVersion      string              `json:"contentVersion"`
	PanelizationSummary PanelizationSummary `json:"panelizationSummary"`
	ImageLinks          ImageLinks          `json:"imageLinks"`
	Language            string              `json:"language"`
	PreviewLink         string              `json:"previewLink"`
	infoLink            string              `json:"infoLink"`
	CanonicalVolumeLink string              `json:"canonicalVolumeLink"`
}

type PanelizationSummary struct {
	ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
	ContainsImageBubbles bool `json:"containsImageBubbles"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type IndustryIdentifiers []struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
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

type SearchResponse struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
}

type SearchTest struct {
	ItemCount int64 `json:"itemCount"`
}

func (app *Config) Search(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("q")
	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=" + params)
	if err != nil {
		log.Println(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sq := SearchQuery{}

	err = json.Unmarshal(respBody, &sq)
	if err != nil {
		log.Println(err)
	}

	_ = app.jsonWrite(w, http.StatusOK, sq)
}

func (app *Config) SearchTest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=dataintensive")
	if err != nil {
		log.Println(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sq := SearchQuery{}

	err = json.Unmarshal(respBody, &sq)
	if err != nil {
		log.Println(err)
	}

	st := SearchTest{
		ItemCount: sq.TotalItems,
	}

	_ = app.jsonWrite(w, http.StatusOK, st)
}
