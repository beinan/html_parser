package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {

	// // Extract all the entries
	// doc.Find("article").Each(func(i int, s *goquery.Selection) {
	// 	entry := s.Find("a")
	// 	title := entry.Text()
	// 	link, _ := entry.Attr("href")
	// 	date := s.Find(".b6")
	// 	dateText := date.Text()
	// 	fmt.Printf("%d; \"%s\"; \"%s\"; https://www.databricks.com%s\n", i+1, dateText, title, link)
	// })

	// html := fetch(`https://resources.snowflake.com/`)

	// doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Extract all the entries
	// doc.Find("article").Each(func(i int, s *goquery.Selection) {
	// 	link, _ := s.Find("a").Attr("href")
	// 	title := s.Find("h1").Text()
	// 	datetime, _ := s.Find(".js-readable-timestamp").Attr("datetime")
	// 	fmt.Printf("%d; \"%s\"; \"%s\"; \"%s\"; \"%s\"\n", i, strings.Split(link, "/")[3], datetime, title, link)
	// })

	html := fetch(`https://www.databricks.com/dataaisummit/sessions`)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		log.Fatal(err)
	}

	// Extract all the entries
	doc.Find("a.gendaItem").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".css-1vov0py").Text()
		title := s.Find("h1").Text()
		datetime, _ := s.Find(".js-readable-timestamp").Attr("datetime")
		fmt.Printf("%d; \"%s\"; \"%s\"; \"%s\"; \"%s\"\n", i, strings.Split(link, "/")[3], datetime, title, link)
	})

}

func databrick() {

	html := fetch(`https://www.databricks.com/blog/`)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		log.Fatal(err)
	}

	// Extract all the entries
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		entry := s.Find("a")
		title := entry.Text()
		link, _ := entry.Attr("href")
		if strings.HasPrefix(link, "/blog/category/") && len(strings.Split(link, "/")) > 4 {
			//fmt.Printf("%d; \"%s\"; https://www.databricks.com%s\n", i+1, title, link)
			extractBlogs("https://www.databricks.com"+link, title)
			extractBlogs("https://www.databricks.com"+link+"/page/2", title)
		}

	})
}

func extractBlogs(uri string, category string) {
	html := fetch(uri)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		log.Fatal(err)
	}
	// Extract all the entries
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		entry := s.Find("a")
		title := entry.Text()
		link, _ := entry.Attr("href")
		date := s.Find(".b6")
		dateText := date.Text()
		if strings.Contains(link, "/2023/") {
			fmt.Printf("\"%s\"; \"%s\"; \"%s\"; https://www.databricks.com%s; %s\n", category, dateText, title, link, uri)
		}
	})
}

func fetch(uri string) string {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	// ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	// defer func() {
	// 	cancel()
	// 	// Prevent Chromium processes from hanging
	// 	if _, err := exec.Command("pkill", "-g", "0", "Chromium").Output(); err != nil {
	// 		log.Println("[warn] Failed to kill Chromium processes", err)
	// 	}
	// }()

	// run task list
	var res string
	// var sel interface{}
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(uri),
		//chromedp.Navigate(`https://www.databricks.com/resources`),
		//chromedp.WaitVisible("main"),
		//chromedp.OuterHTML(sel, &res, chromedp.ByJSPath),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#uf-lazy-loader-load-more`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second*2),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
