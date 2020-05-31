package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

// TODO: check against best practices: templates, dir structure

func main() {
	// set max limit on processes
	// measure performance
	runtime.GOMAXPROCS(100)

	sourceURL := parseInput()
	log.Println("Source Sitemap: ", sourceURL)

	var si = &SitemapIndex{}

	// TODO: Write to file in the end
	// TODO: Show in a data table
	// TODO: Break into separate files
	sitemapURLs := si.getSitemapUrls(sourceURL)

	// with a type alias I do not need constructor function, still need make to initialize map
	// doing this to reduce allocations during adding data
	var sm = make(Sitemaps, len(sitemapURLs))
	sm.parseSitemaps(sitemapURLs)

	log.Println("Number of links gathered: ", len(sm))
	//log.Println(sm)
}

func getSitemapData(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}
