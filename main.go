package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const domain = "https://qiita.com"

func main() {
	reqUrl := domain + "/tags/go/items"
	existsNextPage := true

	for existsNextPage {
		resp, err := http.Get(reqUrl)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".tsf-ArticleList_view").Each(func(_ int, s *goquery.Selection) {
			articleTitle := s.Find(".tsf-ArticleBody a").Text()

			fmt.Println(articleTitle)
		})

		nextPagePath, exists := doc.Find(".st-Pager_next a").Attr("href")
		fmt.Println(nextPagePath)
		if exists {
			reqUrl = domain + nextPagePath
		} else {
			existsNextPage = false
		}
	}

	os.Exit(1)
}
