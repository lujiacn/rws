package rwsStruct

import "encoding/xml"

type SiteRef struct {
	XMLName     xml.Name `xml:"SiteRef"`
	LocationOID string   `xml:"LocationOID,attr"`
}
