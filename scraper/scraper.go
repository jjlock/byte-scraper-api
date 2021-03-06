// Package scraper implements a framework for scraping the byte website.
package scraper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ByteBaseURL is the base URL of the byte website.
// It is used by a Scraper to construct URLs for scraping data.
const ByteBaseURL = "https://byte.co"

// Scraper represents the instance used for scraping the byte website.
type Scraper struct {
	client *http.Client
}

// NewScraper creates a new Scraper instance.
func NewScraper() *Scraper {
	return &Scraper{client: &http.Client{Timeout: 10 * time.Second}}
}

// NewCustomScraper creates a new Scraper instance with the given http.Client.
func NewCustomScraper(c *http.Client) *Scraper {
	return &Scraper{client: c}
}

// get sends a GET request to the specifed url.
// *RequestError is returned on a non-200 response, otherwise it returns
// any error returned from sending the request or parsing the response.
func (s *Scraper) get(url string) (*goquery.Document, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, &RequestError{
			StatusCode: resp.StatusCode,
			Message:    "byte.co responded with HTTP status: " + resp.Status,
		}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	return doc, nil
}
