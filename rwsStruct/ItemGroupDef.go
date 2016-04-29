package rwsStruct

import "encoding/xml"

type ItemGroupDef struct {
	XMLName   xml.Name `xml:"ItemGroupDef"`
	OID       string   `xml:"OID,attr"`
	Name      string   `xml:"Name,attr"`
	Repeating string   `xml:"Repeating,attr"`
	ItemRef   []ItemRef
}
