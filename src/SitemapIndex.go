package main

import (
	"encoding/xml"
)

// combined struct just to make it simple
type SitemapIndex struct {
	// must be capital so they're exposed to xml encoding package
	Locations []string `xml:"sitemap>loc"`
}

func (si *SitemapIndex) getSitemapUrls(url string) []string {
	mainSitemapBytes := getSitemapData(url)
	_ = xml.Unmarshal(mainSitemapBytes, &si)
	return si.Locations
}
