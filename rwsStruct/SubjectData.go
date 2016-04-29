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
