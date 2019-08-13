package rwsStruct

type ClinicalData struct {
	//XMLName              xml.Name    `xml:"ClinicalData"`
	StudyOID             string      `xml:"StudyOID,attr"`
	MetaDataVersionOID   string      `xml:"MetaDataVersionOID,attr"`
	AuditSubCategoryName string      `xml:",mdsol:AuditSubCategoryName,attr"`
	SubjectData          SubjectData //`xml:"ClinicalData>SubjectData"`
}
