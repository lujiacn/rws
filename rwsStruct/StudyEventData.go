package rwsStruct

type StudyEventData struct {
	//XMLName             xml.Name `xml:"StudyEventData"`
	StudyEventOID       string     `xml:"StudyEventOID,attr"`
	StudyEventRepeatKey string     `xml:"StudyEventRepeatKey,attr"`
	InstanceId          string     `xml:",mdsol:InstanceId,attr"`
	FormData            []FormData // can have multiple forms
}
