package rwsStruct

import "encoding/xml"

type ItemData struct {
	XMLName xml.Name `xml:"ItemData"`
	ItemOID string   `xml:"ItemOID,attr"`
	Value   string   `xml:"Value,attr"`
}
