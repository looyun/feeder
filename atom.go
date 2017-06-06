package feeder

import (
	"encoding/xml"
	"io"
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

func ParseAtom(feed io.Reader) (*Feed, error) {
	af := AtomFeed{}
	d := xml.NewDecoder(feed)
	d.Decode(&af)
	result := new(Feed)
	result.Title = af.Title
	result.Author = ParsePerson(&af.Author)

	result.Updated = af.Updated
	result.Link = af.Link
	result.Category = af.Category

	result.Contributor = ParsePerson(&af.Contributor)
	// result.Icon.URL = af.Icon.URL
	// result.Logo.URL = af.Logo.URL
	result.Copyights = af.Rights
	result.Subtitle = af.Subtitle
	result.Generator = af.Generator

	result.Items = ParseEntries(af.Entries)

	return result, nil
}

func ParseEntries(ae []AtomEntry) []*Item {
	items := []*Item{}
	for _, v := range ae {
		item := new(Item)
		item.Author = ParsePerson(&v.Author)
		item.Content = v.Content
		item.Link = v.Link
		item.Updated = v.Updated
		item.Summary = v.Summary
		item.Category = v.Category
		item.Published = v.Published
		// item.Source = v.Source
		item.Contributor = ParsePerson(&v.Contributor)
		items = append(items, item)
	}
	return items
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
