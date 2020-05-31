package main

import (
	"flag"
	"log"
	"net/url"
	"os"
)

func parseInput() string {
	sourceURL := flag.String(
		"sourceURL",
		"https://www.infidigit.com/sitemap_index.xml",
		"Source url of main sitemap")
	flag.Parse()

	if !isValidUrl(*sourceURL) {
		log.Printf("Not a valid url: [%s]\n", *sourceURL)
		os.Exit(1)
	}

	return *sourceURL
}

// isValidUrl tests a string to determine if it is a well-structured url or not.
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
