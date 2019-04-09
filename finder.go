package macfinder

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	apiEndpoint = "https://www.apple.com/fr/shop/refurbished/mac/macbook-pro"

	scriptPrefix = "window.REFURB_GRID_BOOTSTRAP = "
	scriptSuffix = ";"
)

var (
	ErrInvalidRawData = errors.New("invalid data found in the script tag")
)

type response struct {
	Tiles []struct {
		Filters struct {
			Dimensions Specs `json:"dimensions"`
		} `json:"filters"`
		ProductDetailsUrl string `json:"productDetailsUrl"`
	} `json:"tiles"`
}

func FetchAvailable() ([]*Model, error) {
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	scriptRawText := strings.TrimSpace(doc.Find("div[role=main] script").Text())
	if !strings.HasPrefix(scriptRawText, scriptPrefix) || !strings.HasSuffix(scriptRawText, scriptSuffix) {
		return nil, ErrInvalidRawData
	}

	jsonData := scriptRawText[len(scriptPrefix):len(scriptRawText) - len(scriptSuffix)]

	var data response
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	available := make([]*Model, len(data.Tiles))
	for i, t := range data.Tiles {
		available[i] = &Model{t.Filters.Dimensions, t.ProductDetailsUrl}
	}

	return available, nil
}

func FindModel(specs Specs) (*Model, error) {
	available, err := FetchAvailable()
	if err != nil {
		return nil, err
	}

	for _, m := range available {
		if specs.equal(m.Specs) {
			return m, nil
		}
	}

	return nil, nil
}