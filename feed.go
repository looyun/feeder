package feeder

import (
	"time"
)

type Feed struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Updated     *time.Time `json:"updated,omitempty"`
	Author      *Person    `json:"author,omitempty"`
	Link        string     `json:"link,omitempty"`
	Category    string     `json:"category,omitempty"`
	Contributor *Person    `json:"contributor,omitempty"`
	Generator   string     `json:"generator,omitempty"`
	Icon        *Image     `json:"icon,omitempty"`
	Logo        *Image     `json:"logo,omitempty"`
	Rights      string     `json:"rights,omitempty"`
	Subtitle    string     `json:"subtitle,omitempty"`
	Entries     []*Entry   `json:"entries"`
}

type Entry struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Updated     *time.Time `json:"updated,omitempty"`
	Author      *Person    `json:"author,omitempty"`
	Content     string     `json:"content,omitempty"`
	Link        string     `json:"link,omitempty"`
	Summary     string     `json:"summary,omitempty"`
	Category    string     `json:"category,omitempty"`
	Contributor *Person    `json:"contributor,omitempty"`
	Published   *time.Time `json:"published,omitempty"`
	Source      *Feed      `json:"source,omitempty"`
}

type Person struct {
	Name  string `json:"name,omitempty"`
	URI   string `json:"uri,omitempty"`
	Email string `json:"email,omitempty"`
}

type Image struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type Link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Type   string `json:"type,omitempty"`
	Length uint   `json:"length,omitempty"`
}
