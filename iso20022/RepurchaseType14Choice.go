package iso20022

// Choice of format for the repurchase transaction type information.
type RepurchaseType14Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType2Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType14Choice) SetCode(value string) {
	r.Code = (*RepurchaseType2Code)(&value)
}

func (r *RepurchaseType14Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
