package rwsStruct

import "encoding/xml"

type Odm struct {
	XMLName      xml.Name       `xml:"ODM"`
	ClinicalData []ClinicalData `xml:"ClinicalData"`
}
