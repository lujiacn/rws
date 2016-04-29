package rwsStruct

type ItemGroupRef struct {
	XMLName      string `xml:"ItemGroupRef"`
	ItemGroupOID string `xml:"ItemGroupOID,attr"`
	Mandatory    string `xml:"Mandatory,attr"`
}
