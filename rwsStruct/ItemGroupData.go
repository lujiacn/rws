package rwsStruct

import "encoding/xml"

type ItemGroupData struct {
	XMLName             xml.Name `xml:"ItemGroupData"`
	ItemGroupOID        string   `xml:"ItemGroupOID,attr"`
	ItemGroupRepeateKey string   `xml:"ItemGroupRepeateKey,attr"`
	ItemData            []ItemData
}
