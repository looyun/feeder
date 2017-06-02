package feeder

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type AtomFeed struct {
	ID          string      `xml:"id"`
	Title       string      `xml:"title"`
	Updated     *time.Time  `xml:"updated"`
	Author      AtomPerson  `xml:"author"`
	Link        string      `xml:"link"`
	Category    string      `xml:"category"`
	Contributor AtomPerson  `xml:"contributor"`
	Generator   string      `xml:"generator"`
	Icon        AtomImage   `xml:"icon"`
	Logo        AtomImage   `xml:"logo"`
	Rights      string      `xml:"rights"`
	Subtitle    string      `xml:"subtitle"`
	Entries     []AtomEntry `xml:"entry"`
}

type AtomEntry struct {
	ID          string     `xml:"id"`
	Title       string     `xml:"title"`
	Updated     *time.Time `xml:"updated"`
	Author      AtomPerson `xml:"author"`
	Content     string     `xml:"content"`
	Link        string     `xml:"link"`
	Summary     string     `xml:"summary"`
	Category    string     `xml:"category"`
	Contributor AtomPerson `xml:"contributor"`
	Published   *time.Time `xml:"published"`
	Source      AtomFeed   `xml:"source"`
}

type AtomPerson struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
}

type AtomImage struct {
	URL  string `xml:"url"`
	Name string `xml:"name"`
}

type AtomLink struct {
	Href   string `xml:"href,attr"`
	Rel    string `xml:"rel,attr"`
	Type   string `xml:"type,attr"`
	Length uint   `xml:"length,attr"`
}

type Parser struct {
}

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
	af := AtomFeed{}
	d := xml.NewDecoder(feed)
	d.Decode(&af)
	result.Title = af.Title
	result.Author = ParsePerson(&af.Author)

	result.Updated = af.Updated
	result.Link = af.Link
	result.Category = af.Category

	result.Contributor = ParsePerson(&af.Contributor)
	// result.Icon.URL = af.Icon.URL
	// result.Logo.URL = af.Logo.URL
	result.Rights = af.Rights
	result.Subtitle = af.Subtitle
	result.Generator = af.Generator

	result.Entries = ParseEntries(af.Entries)

	return result, nil
}

func ParseEntries(ae []AtomEntry) []*Entry {
	entries := []*Entry{}
	for _, v := range ae {
		entry := new(Entry)
		entry.Author = ParsePerson(&v.Author)
		entry.Content = v.Content
		entry.Link = v.Link
		entry.Updated = v.Updated
		entry.Summary = v.Summary
		entry.Category = v.Category
		entry.Published = v.Published
		// entry.Source = v.Source
		entry.Contributor = ParsePerson(&v.Contributor)
		entries = append(entries, entry)
	}
	return entries
}

func ParsePerson(ap *AtomPerson) *Person {
	person := new(Person)
	person.Name = ap.Name
	person.URI = ap.URI
	person.Email = ap.Email
	return person
}

func ParseImage(ai *AtomImage) *Image {
	image := new(Image)
	image.Name = ai.Name
	image.URL = ai.URL
	return image
}
