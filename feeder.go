package feeder

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func ParseURL(URL string) (*Feed, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return Parse(resp.Body)
}

func Parse(feed io.Reader) (*Feed, error) {
	result := new(Feed)
	var err error

	var buf bytes.Buffer
	tee := io.TeeReader(feed, &buf)

	ft := DetectFeedType(tee)

	r := io.MultiReader(&buf, feed)

	switch ft {
	case "feed":
		// fmt.Println("Atom FeedType")
		result, err = ParseAtom(r)
		return result, err
	case "rss":
		// fmt.Println("RSS FeedType")
		result, err := ParseRSS(r)
		return result, err
	}
	return result, err
}

func DetectFeedType(feed io.Reader) string {
	d := xml.NewDecoder(feed)
	for {
		token, err := d.Token()
		if err != nil {
			fmt.Println(err)
			break
		}
		tokentype, ok := token.(xml.StartElement)
		if ok {
			return tokentype.Name.Local
		}
	}
	return ""
}
