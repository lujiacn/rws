package rwsStruct

type ItemData struct {
	//XMLName         xml.Name `xml:"ItemData"`
	ItemOID         string      `xml:"ItemOID,attr"`
	TransactionType string      `xml:"TransactionType,attr"`
	Value           string      `xml:"Value,attr"`
	IsNull          string      `xml:"IsNull,attr"`
	AuditRecord     AuditRecord // one time one audit record
	Query           Query       // one item one query
}

//<ItemData ItemOID="SUB_ID.SYSDAT" TransactionType="Upsert" Value="17 JUL 2019">
