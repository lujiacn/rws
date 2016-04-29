package rwsStruct

import "encoding/xml"

type StudyEventData struct {
	XMLName             xml.Name `xml:"StudyEventData"`
	StudyEventOID       string   `xml:"StudyEventOID,attr"`
	StudyEventRepeatKey string   `xml:"StudyEventRepeatKey,attr"`
	FormData            FormData
}
