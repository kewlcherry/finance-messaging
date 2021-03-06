package iso20022

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails25 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType7Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition5Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType8Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails25) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails25) AddSecuritiesTransactionType() *SecuritiesTransactionType7Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType7Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails25) AddSettlementTransactionCondition() *SettlementTransactionCondition5Choice {
	newValue := new(SettlementTransactionCondition5Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails25) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails25) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails25) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails25) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails25) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails25) AddRepurchaseType() *RepurchaseType8Choice {
	s.RepurchaseType = new(RepurchaseType8Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails25) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails25) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}
