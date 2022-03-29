package main

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Book struct {
	Title     string    `xml:"http://www.example.com/ title"`
	Published time.Time `xml:"http://www.example.com/ published"`
	Author    string    `xml:"http://www.example.com/ author"`
}

func (t *Book) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Book
	var layout struct {
		*T
		Published *xsdDate `xml:"http://www.example.com/ published"`
	}
	layout.T = (*T)(t)
	layout.Published = (*xsdDate)(&layout.T.Published)
	return e.EncodeElement(layout, start)
}
func (t *Book) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Book
	var overlay struct {
		*T
		Published *xsdDate `xml:"http://www.example.com/ published"`
	}
	overlay.T = (*T)(t)
	overlay.Published = (*xsdDate)(&overlay.T.Published)
	return d.DecodeElement(&overlay, &start)
}

type Library struct {
	Book []Book `xml:"http://www.example.com/ book"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}
