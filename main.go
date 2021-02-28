package main

import (
	"encoding/json"
	"fmt"
	"gotraining/order"
	"os"
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

	//c.Visit("https://www.ubereats.com/tw")
	c.Visit("https://www.ubereats.com/tw/taipei/food-delivery/%E8%AA%A0%E8%A8%98%E5%8E%9F%E6%B1%81%E6%8E%92%E9%AA%A8%E6%B9%AF-%E9%80%9A%E5%8C%96%E5%BA%97/dyF-ZHWCTYu7_bzkhuRWyg")
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
		fmt.Println(e)
		// firstChild := e.DOM.Children().First()
		//TODO:還在特徵測試
		header := e.DOM.Find(".bs").Text()
		text1 := e.DOM.Find(".h3.gx.h4.gy.al.h5").Text()
		text2 := e.DOM.Find(".fc.ag.bu.h0").Text()

		// header := e.DOM.Find("h4").Text()
		// text1 := e.DOM.Find("h4+div").Text()
		// text2 := e.DOM.Find("h4+div+div").Text()
		fmt.Println(text1)
		orderData.Name = header
		if len(text2) > 0 {
			orderData.Description = text1
			orderData.Price = text2
		} else {
			orderData.Description = ""
			orderData.Price = text1
		}

		//fmt.Println(orderData)
		writeFile(parseJsonl(orderData))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(subPageURL)
}

func parseJsonl(order order.Myorder) string {
	out, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func writeFile(words string) {
	file, err := os.OpenFile("appData/orderData.jsonl", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 6044)
	if err != nil {
		fmt.Println("開檔錯誤!")
		panic(err)
	}

	file.WriteString(words + "\n")
	defer file.Close()
}
