package sitemap

import (
	"encoding/xml"
)

type urlset struct {
	XMLName xml.Name   `xml:"urlset"`
	Xmlns   string     `xml:"xmlns,attr"`
	URL     []location `xml:"url"`
}

type location struct {
	Loc string `xml:"loc"`
}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

//GenerateSitemap takes a slice of locations and returns a sitemap in standard sitemap protocol (XML)
func GenerateSitemap(locations []string) (string, error) {
	var sm urlset
	sm.Xmlns = xmlns
	sm.URL = make([]location, 0)
	for _, l := range locations {
		loc := location{l}
		sm.URL = append(sm.URL, loc)
	}

	out, err := xml.MarshalIndent(sm, "", "")
	if err != nil {
		return "", err
	}

	return xml.Header + string(out), nil
}
