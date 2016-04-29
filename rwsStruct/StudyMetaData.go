package rwsStruct

type StudyMetaData struct {
	XMLName         string `xml:"Study"`
	OID             string `xml:"OID,attr"`
	GlobalVariables GlobalVariables
	MetaDataVersion []MetaDataVersion
}
