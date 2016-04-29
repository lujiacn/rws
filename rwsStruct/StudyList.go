//url link format http://.../RaveWebServices/studies
package rwsStruct

import "encoding/xml"

type StudyList struct {
	XMLName xml.Name `xml:"ODM"`
	Studies []Study  `xml:"Study"`
}
