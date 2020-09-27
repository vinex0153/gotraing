package main

import (
	"fmt"
	"gotraining/order"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		//fmt.Println(e.Attr("href"))
		var int = strings.Index(e.Attr("href"), "food-delivery")
		if int > -1 {
			//e.Request.Visit(e.Attr("href"))
			subPage(e.Request.AbsoluteURL(e.Attr("href")))
			// fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))

		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.ubereats.com/tw")
}

func subPage(subPageURL string) {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		colly.AllowedDomains("www.ubereats.com"),
	)

	// Find and visit all links
	c.OnHTML("li li", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.Find("h4").Text())

		orderData := order.Myorder{Name: "A", Description: "B", Price: "C"}
		header := e.DOM.Find("h4").Text()
		text1 := e.DOM.Find("h4+div").Text()
		text2 := e.DOM.Find("h4+div+div").Text()
		fmt.Println(text2)
		orderData.Name = header
		if len(text2) > 0 {
			orderData.Description = text1
			orderData.Price = text2
		} else {
			orderData.Description = ""
			orderData.Price = text1
		}

		fmt.Println(orderData)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(subPageURL)
}
