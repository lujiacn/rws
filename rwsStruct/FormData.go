package rwsStruct

type FormData struct {
	//XMLName        xml.Name `xml:"FormData"`
	FormOID       string          `xml:"FormOID,attr"`
	FormRepeatKey string          `xml:"FormRepeatKey,attr"`
	DataPageId    string          `xml:",mdsol:DataPageId,attr"`
	ItemGroupData []ItemGroupData // may have multiple group
}
