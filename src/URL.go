package main

import "fmt"

type URL struct {
	Location     string `xml:"loc"`
	LastModified string `xml:"lastmod"`
}

func (p URL) String() string {
	return fmt.Sprintf("URL: %v, LastModified: %v\n", p.Location, p.LastModified)
}
