package rss

import "encoding/xml"

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName         xml.Name `xml:"channel"`
	Title           string   `xml:"title"`
	Description     string   `xml:"description"`
	Link            string   `xml:"link"`
	Language        string   `xml:"language"`
	PublicationDate string   `xml:"pubDate"`
	LastBuildDate   string   `xml:"lastBuildDate"`
	TTL             int      `xml:"ttl"`
	Items           []Item   `xml:"item"`
}

type Item struct {
	XMLName         xml.Name `xml:"item"`
	Title           string   `xml:"title"`
	Description     string   `xml:"description"`
	Link            string   `xml:"link"`
	PublicationDate string   `xml:"pubDate"`
	Categories      []string `xml:"category"`
}

func NewRSSFrom(data []byte) (rss *RSS, err error) {
	rss = new(RSS)
	err = xml.Unmarshal(data, rss)
	return
}
