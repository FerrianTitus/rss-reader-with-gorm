package model

import (
	"encoding/xml"
	"time"
)

type Channel struct {
	XMLName xml.Name  `xml:"channel"`
	Items   []RSSItem `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type RSSItem struct {
	XMLNAME     xml.Name   `xml:"item"`
	ID          int64      `gorm: "primary_key" xml:"-"`
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Link        string     `xml:"link"`
	CreatedAt   time.Time  `xml:"-"`
	UpdatedAt   time.Time  `xml:"-"`
	DeletedAt   *time.Time `xml:"-"`
}
