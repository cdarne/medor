package rss

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`

	// required
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`

	// optional
	Language        string     `xml:"language"`
	PublicationDate RFC822Time `xml:"pubDate"`
	LastBuildDate   RFC822Time `xml:"lastBuildDate"`
	TTL             int        `xml:"ttl"`
	Copyright       string     `xml:"copyright"`
	ManagingEditor  string     `xml:"managingEditor"`
	WebMaster       string     `xml:"webMaster"`

	Items []Item `xml:"item"`
}

type Item struct {
	XMLName         xml.Name   `xml:"item"`
	Title           string     `xml:"title"`
	Description     string     `xml:"description"`
	Link            string     `xml:"link"`
	PublicationDate RFC822Time `xml:"pubDate"`
	Categories      []string   `xml:"category"`
}

func NewRSSFrom(data []byte) (rss *RSS, err error) {
	rss = new(RSS)
	err = xml.Unmarshal(data, rss)
	return
}

type RFC822Time time.Time

var timeFormats = [...]string{time.RFC822, time.RFC822Z, time.RFC1123, time.RFC1123Z}

func (t *RFC822Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var str string
	var parsedTime time.Time

	err = d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	for _, timeFormat := range timeFormats {
		parsedTime, err = time.Parse(timeFormat, str)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}

	*t = (RFC822Time)(parsedTime)
	return nil
}

func (t RFC822Time) String() string {
	return (time.Time)(t).String()
}
