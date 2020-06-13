package main

import (
	"flag"
	"fmt"

	"github.com/vncsb/sitemap"
)

var url = flag.String("url", "http://www.basicwebsiteexample.com/", "Defines the URL where the mapper will start. Example: http://www.basicwebsiteexample.com/")

func main() {
	flag.Parse()

	s, err := sitemap.Map(*url)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
