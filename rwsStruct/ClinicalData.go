package rwsStruct

import "encoding/xml"

type ClinicalData struct {
	XMLName            xml.Name    `xml:"ClinicalData"`
	StudyOID           string      `xml:"StudyOID,attr"`
	MetaDataVersionOID string      `xml:"MetaDataVersionOID,attr"`
	SubjectData        SubjectData `xml:"SubjectData"`
}
