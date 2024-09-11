package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/gocolly/colly"
)

type Product struct {
	Url   string
	Image string
	Name  string
	Price string
}

func main() {
	domain := "www.scrapingcourse.com"
	url := fmt.Sprintf("https://%s/ecommerce/", domain)

	var products []Product

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		products = append(products, Product{
			e.ChildAttr("a", "href"),
			e.ChildAttr("img", "src"),
			e.ChildText(".product-name"),
			e.ChildText(".price"),
		})

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped: ", r.Request.URL)

		productsJson, _ := json.MarshalIndent(products, "", "\t")
		os.WriteFile("products.json", productsJson, 0644)
	})

	err := c.Visit(url)

	// This error is necessary to detect error prior to the visit
	if err != nil {
		fmt.Println("Visit Error:", err)
	}
}
