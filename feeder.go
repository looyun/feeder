package feeder

import (
	"bytes"
	"encoding/xml"
	"errors"
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

	ft, err := DetectFeedType(tee)
	if err != nil {
		return nil, err
	}

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
	default:
		return nil, errors.New("Unknown FeedType.")
	}
	return result, err
}

func DetectFeedType(feed io.Reader) (string, error) {
	d := xml.NewDecoder(feed)
	for {
		token, err := d.Token()
		if err != nil {
			return "", err
		}
		tokentype, ok := token.(xml.StartElement)
		if ok {
			return tokentype.Name.Local, nil
		}
	}
	return "unknown", errors.New("Unknown content.")
}
