package rwsStruct

import "encoding/xml"

type MetaDataVersion struct {
	XMLName      xml.Name `xml:"MetaDataVersion"`
	OID          string   `xml:OID,attr`
	Name         string   `xml:"Name,attr"`
	FormDef      []FormDef
	ItemGroupDef []ItemGroupDef
	ItemDef      []ItemDef
}
