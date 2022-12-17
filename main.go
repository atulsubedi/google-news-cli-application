package main

import(
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
	"net/http"

)


type RSS struct {
	XMLName		xml.Name 
	Channel		*Channel 
}

type Channel struct {
	Title 		string	
	ItemList	[]Item 
}

type Item struct {
	Title 			string
	Link 			string
	Traffic 		string
	NewsItems		[]News
}

type News struct {
	HeadLine			string
	HeadLineLink		string
}

func main() {
	readGoogleTrends
}

func readGoogleTrends(){
	getGoogleTrends
}

func getGoogleTrends(){

}