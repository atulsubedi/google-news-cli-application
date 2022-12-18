package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	err :=	xml.Unmarshal(data, &r)

	if err != nil{
		fmt.Println("errro:", err)
		os.Exit(1)
	}
}

func getGoogleNews() *http.Response {
	resp, err := http.Get("https://news.google.com/rss?hl=en-US&gl=US&ceid=US:en")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readGoogleNews() []byte {
	resp := getGoogleNews()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
