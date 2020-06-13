package sitemap

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/vncsb/linkparser"
)

type page struct {
	links    []linkparser.Link
	domain   string
	location string
	scheme   string
}

//Map takes a link and returns a sitemap in standard sitemap protocol (XML)
func Map(link string) (string, error) {
	locations := map[string]bool{}
	crawl(link, locations)

	sitemap, err := GenerateSitemap(getKeysFromMap(locations))
	if err != nil {
		return "", err
	}

	return sitemap, nil
}

func parsePage(link string) (page, error) {
	resp, err := http.Get(link)
	if err != nil {
		return page{}, err
	}

	defer resp.Body.Close()
	links, err := linkparser.Parse(resp.Body)
	if err != nil {
		return page{}, err
	}

	u := resp.Request.URL

	return page{
		links:    links,
		domain:   u.Hostname(),
		location: u.String(),
		scheme:   u.Scheme,
	}, nil
}

func crawl(link string, foundLocations map[string]bool) {
	page, err := parsePage(link)
	if err != nil {
		return
	}

	foundLocations[page.location] = true
	fmt.Println("Found: " + page.location)
	domain := page.domain
	for _, l := range page.links {
		location := formatLink(l.Href, domain, page.scheme)
		_, ok := foundLocations[location]
		if isSameDomain(location, domain) && !ok {
			crawl(location, foundLocations)
		}
	}
}

func formatLink(link string, domain string, scheme string) string {
	if strings.HasPrefix(link, "#") {
		return "/"
	}
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = scheme + "://" + domain + link
	}
	return link
}

func isSameDomain(link string, domain string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	return u.Hostname() == domain
}

func getKeysFromMap(m map[string]bool) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
