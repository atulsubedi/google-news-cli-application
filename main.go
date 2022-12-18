package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

func main() {
	var r RSS
	data := readGoogleNews()
}

func getGoogleNews() *http.Response {
	resp, err := http.Get("https://news.google.com/rss?hl=en-US&gl=US&ceid=US:en")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readGoogleNews() []byte {
	getGoogleNews
}
