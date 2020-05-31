package main

import (
	"encoding/xml"
	"log"
)

// This has to be struct. Can't be a type alias due to the tag being only on property level
type URLs struct {
	URL []URL `xml:"url"`
}

// avoid pointer receivers in go routines, these can cause race conditions
func (p URLs) parseSitemapURL(url string, ch chan<- map[string]URLs) {
	log.Println("Parsing: ", url)
	pages := getSitemapData(url)
	_ = xml.Unmarshal(pages, &p)
	m := make(map[string]URLs)
	m[url] = p

	// using select only to indicate if channel gets full
	select {
	case ch <- m:
	default:
		log.Fatalln("Channel full. Discarding value")
	}
}

func (p URLs) String() string {
	var str string
	for _, v := range p.URL {
		str += v.String()
	}
	return str
}
