package rwsStruct

import "encoding/xml"

type FormDef struct {
	XMLName      xml.Name `xml:"FormDef"`
	OID          string   `xml:OID,attr`
	Name         string   `xml:"Name,attr"`
	Repeating    string   `xml:"Repeating,attr"`
	ItemGroupRef ItemGroupRef
}
