package webservice

import "encoding/xml"

type Response struct {
	XMLName xml.Name `xml:"response"`
	Content string `xml:"content"`
	OriginalContent string `xml:"original-content"`
}