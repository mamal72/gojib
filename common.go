package gojib

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getAndParseURL(url string) (*goquery.Document, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func stringToNumber(input string) int {
	normalizedInput := strings.Replace(input, ",", "", -1)
	normalized, _ := strconv.Atoi(normalizedInput)
	return normalized
}
