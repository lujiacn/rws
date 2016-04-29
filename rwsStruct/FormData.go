package rwsStruct

import "encoding/xml"

type FormData struct {
	XMLName        xml.Name `xml:"FormData"`
	FormOID        string   `xml:"FormOID,attr"`
	FormRepeateKey string   `xml:"FormRepeateKey,attr"`
	ItemGroupData  ItemGroupData
}
