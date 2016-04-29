package rwsStruct

import "encoding/xml"

type FormMetaData struct {
	XMLName xml.Name      `xml:"ODM"`
	Study   StudyMetaData `xml:"Study"`
}
