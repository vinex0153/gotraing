package main

import (
	"fmt"

	. "gotraining/gotraing/class/order"

	"github.com/gocolly/colly"
)

// func main() {

// 	//var aaa = []string{}

// 	c := colly.NewCollector(
// 		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
// 	)

// 	// Find and visit all links
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		//e.Request.Visit(e.Attr("href"))
// 		//fmt.Println(e.Attr("href"))
// 		var int = strings.Index(e.Attr("href"), "food-delivery")
// 		if int > -1 {
// 			e.Request.Visit(e.Attr("href"))
// 		}

// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.Visit("https://www.ubereats.com/tw")
// }

// func visitSubPage(subPageUrl string){

// subC:=colly.NewCollector(
// 	colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
// )

// subc.OnHTML("h4", func(e *colly.HTMLElement) {

// })

// subC.Visit()
// }

func main() {

	//var aaa = []string{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// Find and visit all links
	c.OnHTML("li li", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Find("h4").Text())

		orderData := order{Name: "A", Description: "B", Price: "C"}
		header := e.DOM.Find("h4").Text()
		text1 := e.DOM.Find("h4+div").Text()
		text2 := e.DOM.Find("h4+div+div").Text()

		orderData.name = header
		if len(text2) > 0 {
			orderData.description = ""
			orderData.price = text1
		} else {
			orderData.description = text1
			orderData.price = text2
		}

		fmt.Println(orderData)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.ubereats.com/tw/taipei/food-delivery/%E9%BA%A5%E5%91%B3%E7%99%BB-%E5%8C%97%E5%B8%82%E6%B0%91%E7%94%9F%E6%95%A6%E5%8C%96%E5%BA%97/S_DqkrFoQlW7_XpZhu3jMw")
}
