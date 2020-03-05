package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/json"
)

// Registry will store the Country registration URL items
type Registry struct {
	Country string
	URL     string
}

func main() {
	registries := []Registry{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".govspeak .govuk-link", func(e *colly.HTMLElement) {
		reg := Registry{}
		reg.Country = e.Text
		reg.URL = e.Attr("href")
		registries = append(registries, reg)
	})

	c.OnScraped(func(r *colly.Response) { 
		data, err := json.Marshal(registries)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Finished. Here is your data:", string(data))
		}
		
	})

	c.Visit("https://www.gov.uk/government/publications/overseas-registries/overseas-registries")
}
