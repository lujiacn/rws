package rwsStruct

import "encoding/xml"

type ItemDef struct {
	XMLName           xml.Name `xml:"ItemDef"`
	OID               string   `xml:"OID,attr"`
	Name              string   `xml:"Name,attr"`
	DataType          string   `xml:"DataType,attr"`
	Length            string   `xml:"Length,attr"`
	SignificantDigits string   `xml:"SignificantDigits,attr"`
}
