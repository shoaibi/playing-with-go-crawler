package main

import (
	"fmt"
	"sync"
)

type Sitemaps map[string]URLs

// Could also make it a struct but that isn't needed, but then i need constructor function to initialize the map inside
//type Sitemaps struct {
//	maps map[string]URLs
//}
//
//func NewSiteMaps() *Sitemaps  {
//	return &Sitemaps{
//		map[string]URLs{},
//	}
//}

func (s *Sitemaps) produce(sitemapURLs []string) chan map[string]URLs {
	urlsCount := len(sitemapURLs)
	// Intentionally made the channel less than the urls to see if channel gets blocked.
	urlsChan := make(chan map[string]URLs, urlsCount)

	var wg sync.WaitGroup
	wg.Add(urlsCount)

	urls := &URLs{}
	for _, url := range sitemapURLs {
		go func(url string, ch chan<- map[string]URLs) {
			defer wg.Done()
			urls.parseSitemapURL(url, ch)
		}(url, urlsChan)
	}

	go func() {
		wg.Wait()
		close(urlsChan)
	}()

	return urlsChan
}

func (s *Sitemaps) consume(ch chan map[string]URLs) {
	for i := range ch {
		//log.Println("Got a new message on channel: ", i)
		for k, v := range i {
			(*s)[k] = v
		}
	}
}

func (s *Sitemaps) parseSitemaps(sitemapURLs []string) {
	urlsChan := s.produce(sitemapURLs)
	s.consume(urlsChan)
}

// functions can also be on aliases types
func (s Sitemaps) String() string {
	var str string
	for k, v := range s {
		str += fmt.Sprintf("k=%v, v=%v", k, v)
	}
	return str
}
