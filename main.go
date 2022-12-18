package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
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
	Title     string `xml:"title"`
	Link      string `xml:"link"`
}


func main() {
	var r RSS
	readGoogleNews()
}

func getGoogleNews() *http.Response {
		http.Get("")
}

func readGoogleNews() []byte {
	getGoogleNews()
}
