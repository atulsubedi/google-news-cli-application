package main

import(
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
	"net/http"

)


type RSS struct {
	XMLName		xml.Name			`xml:"rss"`			 
	Channel		*Channel			`xml:"channel"` 
}

type Channel struct {
	Title 		string				`xml:"title"`	
	ItemList	[]Item				`xml:"item"` 
}

type Item struct {
	Title 			string			`xml:"title"`	
	Link 			string			`xml:"link"`	
	Traffic 		string			`xml:"approx_traffic"`
	NewsItems		[]News			`xml:"news_item"`
}

type News struct {
	HeadLine		string			`xml:"news_item_title"`
	HeadLineLink		string		`xml:"news_item_url"`
}

func main() {
	readGoogleTrends
}

func readGoogleTrends(){
	getGoogleTrends
}

func getGoogleTrends(){

}