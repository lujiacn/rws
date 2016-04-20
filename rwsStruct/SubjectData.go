package rwsStruct

import "encoding/xml"

type SubjectData struct {
	XMLName                xml.Name `xml:"SubjectData"`
	SubjectKey             string   `xml:"SubjectKey,attr"`
	SubjectKeyType         string   `xml:"mdsol:SubjectKeyType,attr"`
	OverDue                string   `xml:"mdsol:OverDue,attr"`
	Touched                string   `xml:"mdsol:Touched,attr"`
	Empty                  string   `xml:"mdsol:Empty,attr"`
	Incomplete             string   `xml:"mdsol:Incomplete,attr"`
	NonComformant          string   `xml:"mdsol:NonComformant,attr"`
	RequiresSecondPass     string   `xml:"mdsol:RequiresSecondPass,attr"`
	RequiresReconciliation string   `xml:"mdsol:RequiresReconciliation,attr"`
	RequiresVericiation    string   `xml:"mdsol:RequiresVericiation,attr"`
	Verified               string   `xml:"mdsol:Verified,attr"`
	Frozen                 string   `xml:"mdsol:Frozen,attr"`
	SiteRef                SiteRef
	StudyEventData         StudyEventData
	FormData               FormData
}

type SiteRef struct {
	XMLName     xml.Name `xml:"SiteRef"`
	LocationOID string   `xml:"LocationOID,attr"`
}

type StudyEventData struct {
	XMLName             xml.Name `xml:"StudyEventData"`
	StudyEventOID       string   `xml:"StudyEventOID,attr"`
	StudyEventRepeatKey string   `xml:"StudyEventRepeatKey,attr"`
	FormData            FormData
}

type FormData struct {
	XMLName        xml.Name `xml:"FormData"`
	FormOID        string   `xml:"FormOID,attr"`
	FormRepeateKey string   `xml:"FormRepeateKey,attr"`
	ItemGroupData  ItemGroupData
}

type ItemGroupData struct {
	XMLName             xml.Name `xml:"ItemGroupData"`
	ItemGroupOID        string   `xml:"ItemGroupOID,attr"`
	ItemGroupRepeateKey string   `xml:"ItemGroupRepeateKey,attr"`
	ItemData            []ItemData
}

type ItemData struct {
	XMLName xml.Name `xml:"ItemData"`
	ItemOID string   `xml:"ItemOID,attr"`
	Value   string   `xml:"Value,attr"`
}
