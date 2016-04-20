package rwsStruct

import "encoding/xml"

type FormMetaData struct {
	XMLName xml.Name      `xml:"ODM"`
	Study   StudyMetaData `xml:"Study"`
}

type StudyMetaData struct {
	XMLName         string `xml:"Study"`
	OID             string `xml:"OID,attr"`
	GlobalVariables GlobalVariables
	MetaDataVersion []MetaDataVersion
}

type MetaDataVersion struct {
	XMLName      xml.Name `xml:"MetaDataVersion"`
	OID          string   `xml:OID,attr`
	Name         string   `xml:"Name,attr"`
	FormDef      []FormDef
	ItemGroupDef []ItemGroupDef
	ItemDef      []ItemDef
}

type FormDef struct {
	XMLName      string `xml:"FormDef"`
	OID          string `xml:OID,attr`
	Name         string `xml:"Name,attr"`
	Repeating    string `xml:"Repeating,attr"`
	ItemGroupRef ItemGroupRef
}

type ItemGroupRef struct {
	XMLName      string `xml:"ItemGroupRef"`
	ItemGroupOID string `xml:"ItemGroupOID,attr"`
	Mandatory    string `xml:"Mandatory,attr"`
}

type ItemGroupDef struct {
	XMLName   string `xml:"ItemGroupDef"`
	OID       string `xml:"OID,attr"`
	Name      string `xml:"Name,attr"`
	Repeating string `xml:"Repeating,attr"`
	ItemRef   []ItemRef
}

type ItemRef struct {
	XMLName   string `xml:"ItemRef"`
	ItemOID   string `xml:"ItemOID,attr"`
	Mandatory string `xml:"Mandatory,attr"`
}

type ItemDef struct {
	XMLName           string `xml:"ItemDef"`
	OID               string `xml:"OID,attr"`
	Name              string `xml:"Name,attr"`
	DataType          string `xml:"DataType,attr"`
	Length            string `xml:"Length,attr"`
	SignificantDigits string `xml:"SignificantDigits,attr"`
}
