package rwsStruct

import "encoding/xml"

type Subjects struct {
	XMLName      xml.Name       `xml:"ODM"`
	ClinicalData []ClinicalData `xml:"ClinicalData"`
}
