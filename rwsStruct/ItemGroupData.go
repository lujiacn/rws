package rwsStruct

type ItemGroupData struct {
	//XMLName             xml.Name `xml:"ItemGroupData"`
	ItemGroupOID       string `xml:"ItemGroupOID,attr"`
	ItemGroupRepeatKey string `xml:"ItemGroupRepeatKey,attr"`
	RecordId           string `xml:",mdsol:RecordId,attr"`
	ItemData           []ItemData
}
