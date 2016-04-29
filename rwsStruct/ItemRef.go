package rwsStruct

import "encoding/xml"

type ItemRef struct {
	XMLName   xml.Name `xml:"ItemRef"`
	ItemOID   string   `xml:"ItemOID,attr"`
	Mandatory string   `xml:"Mandatory,attr"`
}
