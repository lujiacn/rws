package rwsStruct

import "encoding/xml"

type Study struct {
	XMLName         xml.Name `xml:"Study"`
	OID             string   `xml:"OID,attr"`
	GlobalVariables GlobalVariables
}
