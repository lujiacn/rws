package rwsStruct

type AuditRecord struct {
	//XMLName         xml.Name `xml:"AuditRecord"`
	UserRef         UserRef
	LocationRef     LocationRef
	DateTimeStamp   string `xml:"DateTimeStamp"`
	ReasonForChange string `xml:"ReasonForChange"`
	SourceID        string `xml:"SourceID"`
}

//<UserRef UserOID="Katherinnan2"/>
//<LocationRef LocationOID="TEST SITE-99991"/>
//<DateTimeStamp>2019-07-17T07:13:33</DateTimeStamp>
//<ReasonForChange>Consented</ReasonForChange>
//<SourceID>26754798</SourceID>
