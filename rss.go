package feeder

import (
	"encoding/xml"
	"io"
	"strings"
	"time"
)

// type RSS struct {
// 	Channel RSSChannel `xml:"channel"`
// }

type RSS struct {
	ID             string    `xml:"id"`
	Version        string    `xml:"version,attr"`
	Title          string    `xml:"channel>title"`
	PubDate        string    `xml:"channel>pubDate"`
	LastBuildDate  string    `xml:"channel>lastBuildDate"`
	Description    string    `xml:"channel>description"`
	Language       string    `xml:"channel>language"`
	ManagingEditor string    `xml:"channel>managingEditor"`
	WebMaster      string    `xml:"channel>webMaster"`
	Docs           string    `xml:"channel>docs"`
	Link           string    `xml:"channel>link"`
	Category       string    `xml:"channel>category"`
	Generator      string    `xml:"channel>generator"`
	Copyright      string    `xml:"channel>copyright"`
	TTL            string    `xml:"channel>ttl"`
	Image          RSSImage  `xml:"channel>image"`
	Rating         string    `xml:"channel>rating"`
	TextInput      string    `xml:"channel>testinput"`
	Cloud          string    `xml:"channel>cloud"`
	SkipHours      []int     `xml:"channel>skiphours"`
	SkipDays       []string  `xml:"channel>skipdays"`
	Items          []RSSItem `xml:"channel>item"`
}

type RSSItem struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Comment     string `xml:"comment"`
	Enclosure   string `xml:"enclosure"`
	PubDated    string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Author      string `xml:"author"`
	Link        string `xml:"link"`
	Category    string `xml:"category"`
	Source      RSS    `xml:"source"`
}

type RSSImage struct {
	URL         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Width       string `xml:"width"`
	Height      string `xml:"height"`
	Description string `xml:"description"`
}

func ParseRSS(feed io.Reader) (*Feed, error) {
	RSS := RSS{}
	d := xml.NewDecoder(feed)
	err := d.Decode(&RSS)
	if err != nil {
		return nil, err
	}
	result := new(Feed)
	result.Language = RSS.Language
	result.Title = RSS.Title
	result.Link = RSS.Link
	result.Subtitle = RSS.Description
	author := new(Person)
	author.Email = RSS.WebMaster
	result.Author = author

	generator := new(Person)
	generator.Name = RSS.Generator
	result.Author = author

	result.Updated = ParseDate(RSS.PubDate)
	result.Link = RSS.Link
	result.Category = RSS.Category

	// result.Icon.URL = RSS.Icon.URL
	// result.Logo.URL = RSS.Logo.URL
	result.Copyights = RSS.Copyright
	result.Generator = RSS.Generator

	result.Items = ParseItems(RSS.Items)

	return result, err
}

func ParseItems(RSSItems []RSSItem) []*Item {
	items := []*Item{}
	for _, v := range RSSItems {
		item := new(Item)
		author := new(Person)
		author.Email = v.Author
		item.Author = author
		item.Title = v.Title
		item.Content = v.Description
		item.Link = v.Link
		item.Category = v.Category
		item.Published = ParseDate(v.PubDated)
		// item.Source = v.Source
		items = append(items, item)
	}
	return items
}

func ParseDate(t string) *time.Time {

	then := time.Time{}
	if len(t) >= 25 {
		if strings.HasSuffix(t, "0000") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 +0000", t)
		} else if strings.HasSuffix(t, "GMT") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", t)
		} else if strings.HasSuffix(t, "UTC") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 UTC", t)
		} else if strings.HasSuffix(t, "CST") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 CST", t)
		} else if strings.HasSuffix(t, "0400") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 -0400", t)
		} else if strings.HasSuffix(t, "Z") {
			then, _ = time.Parse(time.RFC3339, t)
		} else if strings.HasSuffix(t, "0800") {
			then, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 +0800", t)
		}
	} else {
		if strings.HasSuffix(t, "0000") {
			then, _ = time.Parse("02 Jan 06 15:04 +0000", t)
		} else if strings.HasSuffix(t, "GMT") {
			then, _ = time.Parse("02 Jan 06 15:04 GMT", t)
		} else if strings.HasSuffix(t, "UTC") {
			then, _ = time.Parse("02 Jan 06 15:04 UTC", t)
		} else if strings.HasSuffix(t, "CST") {
			then, _ = time.Parse("02 Jan 06 15:04 CST", t)
		} else if strings.HasSuffix(t, "0400") {
			then, _ = time.Parse("02 Jan 06 15:04 -0400", t)
		} else if strings.HasSuffix(t, "Z") {
			then, _ = time.Parse(time.RFC3339, t)
		} else if strings.HasSuffix(t, "0800") {
			then, _ = time.Parse("02 Jan 06 15:04 +0800", t)
		}
	}
	return &then
}
