package webpage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadWebPage(webAddr string, sizeChan chan int) {
	fmt.Printf("*** Web page addr: %s ***\n", webAddr)
	response, err := http.Get(webAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	sizeChan <- len(body)
}

type WebPage struct {
	Url  string
	Size int
}

func ReadWebPageStruct(webAddr string, sizeChan chan WebPage) {
	//fmt.Printf("*** Web page addr: %s ***\n", webAddr)
	response, err := http.Get(webAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	sizeChan <- WebPage{Url: webAddr, Size: len(body)}
}
