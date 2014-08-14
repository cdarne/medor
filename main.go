package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/cdarne/medor/rss"
)

func main() {
	logger := log.New(os.Stderr, "example", log.LstdFlags|log.Llongfile)
	data, err := ioutil.ReadFile("lifehacker.rss")
	if err != nil {
		logger.Fatal(err)
	}

	rssObj, err := rss.NewRSSFrom(data)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("RSS: %+v\n", rssObj)
}
