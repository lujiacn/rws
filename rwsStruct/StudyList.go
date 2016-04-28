//url link format http://.../RaveWebServices/studies
package rwsStruct

import "encoding/xml"

type StudyList struct {
	XMLName xml.Name `xml:"ODM"`
	Studies []Study  `xml:"Study"`
}

type Study struct {
	XMLName         xml.Name `xml:"Study"`
	OID             string   `xml:"OID,attr"`
	GlobalVariables GlobalVariables
}

type GlobalVariables struct {
	XMLName          xml.Name `xml:"GlobalVariables"`
	StudyName        string   `xml:"StudyName"`
	StudyDescription string   `xml:"StudyDescription"`
	ProtocolName     string   `xml:"ProtocolName"`
}
