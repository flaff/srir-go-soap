package webservice

import "encoding/xml"

type Request struct {
	XMLName xml.Name `xml:"request"`
	Content	string `xml:"content"`
}