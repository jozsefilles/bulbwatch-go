package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	productUrl :=
		//	"https://lumenet.hu/osram-night-breaker-laser-h7-200-2db/csomag"
		"https://lumenet.hu/tungsram-megalight-ultra-150-h7-duo"
	checkProduct(productUrl)
}

func checkProduct(url string) {
	n := parseNumberOfProductReviews(url)
	fmt.Printf("%v: %v\n", url, n)
}

func parseNumberOfProductReviews(url string) int {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	text := doc.
		Find(".reviews-count").
		Get(0).
		FirstChild.
		Data
	t := strings.Trim(text, "()")
	num, _ := strconv.Atoi(t)
	return num
}
