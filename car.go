package gojib

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CarPrice holds name and price of a car
type CarPrice struct {
	Manufacturer string
	Name         string
	BazaarPrice  int
	CompanyPrice int
}

const (
	iranianCarsPricesURL = "https://www.iranjib.ir/showgroup/45/%D9%82%DB%8C%D9%85%D8%AA-%D8%AE%D9%88%D8%AF%D8%B1%D9%88-%D8%AA%D9%88%D9%84%DB%8C%D8%AF-%D8%AF%D8%A7%D8%AE%D9%84/"
	foreignCarsPricesURL = "https://www.iranjib.ir/showgroup/46/%D9%82%DB%8C%D9%85%D8%AA-%D8%AE%D9%88%D8%AF%D8%B1%D9%88-%D9%88%D8%A7%D8%B1%D8%AF%D8%A7%D8%AA%DB%8C/"

	carTableRowsSelector     = ".items_table:first-of-type tr"
	carManufacturerClassName = "catsection"
	carRowClassName          = ""
	carNameSelector          = "td.name"
	carBazaarPriceSelector   = "td[id$=\"pr\"]:nth-child(1)"
	carCompanyPriceSelector  = "td[id$=\"pr\"]:nth-child(2)"
)

func getCarPricesFromURL(url string) ([]CarPrice, error) {
	doc, err := getAndParseURL(url)
	if err != nil {
		return nil, err
	}
	var currentManufacturer string
	var carsPrices []CarPrice
	doc.Find(carTableRowsSelector).Each(func(_ int, row *goquery.Selection) {
		if row.HasClass(carManufacturerClassName) {
			currentManufacturer = strings.TrimSpace(row.Text())
			return
		}
		if className, _ := row.Attr("class"); className == carRowClassName {
			carsPrices = append(carsPrices, CarPrice{
				Manufacturer: strings.TrimSpace(currentManufacturer),
				Name:         strings.TrimSpace(row.Find(carNameSelector).Text()),
				BazaarPrice:  stringToNumber(strings.TrimSpace(row.Find("td[id]").First().Text())),
				CompanyPrice: stringToNumber(strings.TrimSpace(row.Find("td[id]").Last().Text())),
			})
			return
		}
	})
	return carsPrices, nil
}

// GetIranianCarsPrices return a slice of car prices for Iranian cars
func GetIranianCarsPrices() ([]CarPrice, error) {
	return getCarPricesFromURL(iranianCarsPricesURL)
}

// GetForeignCarsPrices return a slice of car prices for foreign cars
func GetForeignCarsPrices() ([]CarPrice, error) {
	return getCarPricesFromURL(foreignCarsPricesURL)
}
