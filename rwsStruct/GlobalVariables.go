package rwsStruct

import "encoding/xml"

type GlobalVariables struct {
	XMLName          xml.Name `xml:"GlobalVariables"`
	StudyName        string   `xml:"StudyName"`
	StudyDescription string   `xml:"StudyDescription"`
	ProtocolName     string   `xml:"ProtocolName"`
}
