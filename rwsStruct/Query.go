package rwsStruct

type Query struct {
	//XMLName        xml.Name `xml:",mdsol:Query"`
	QueryRepeatKey string `xml:"QueryRepeatKey,attr"`
	Value          string `xml:"Value,attr"`
	Status         string `xml:"Status,attr"`
	Recipient      string `xml:"Recipient,attr"`
}
