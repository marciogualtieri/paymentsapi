package models

/*
Payment entity.
*/
type Payment struct {
	Attributes     Attributes `json:"attributes"`
	ID             string     `json:"id" gorm:"primary_key"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int64      `json:"version"`
}

/*
BeneficiaryParty entity.
*/
type BeneficiaryParty struct {
	ID                uint   `json:"-"`
	AttributesID      uint   `json:"-"`
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int64  `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

/*
DebtorParty entity.
*/
type DebtorParty struct {
	ID                uint   `json:"-"`
	AttributesID      uint   `json:"-"`
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

/*
SponsorParty entity.
*/
type SponsorParty struct {
	ID            uint   `json:"-"`
	AttributesID  uint   `json:"-"`
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
}

/*
Attributes entity.
*/
type Attributes struct {
	ID                   uint               `json:"-"`
	PaymentID            string             `json:"-"`
	Amount               string             `json:"amount"`
	BeneficiaryParty     BeneficiaryParty   `json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `json:"charges_information"`
	Currency             string             `json:"currency"`
	DebtorParty          DebtorParty        `json:"debtor_party"`
	EndToEndReference    string             `json:"end_to_end_reference"`
	Fx                   Fx                 `json:"fx"`
	NumericReference     string             `json:"numeric_reference"`
	PaymentIdentifier    string             `json:"payment_id"`
	PaymentPurpose       string             `json:"payment_purpose"`
	PaymentScheme        string             `json:"payment_scheme"`
	PaymentType          string             `json:"payment_type"`
	ProcessingDate       string             `json:"processing_date"`
	Reference            string             `json:"reference"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type"`
	SchemePaymentType    string             `json:"scheme_payment_type"`
	SponsorParty         SponsorParty       `json:"sponsor_party"`
}

/*
SenderCharges entity.
*/
type SenderCharges struct {
	ID                   uint   `json:"-"`
	ChargesInformationID uint   `json:"-"`
	Amount               string `json:"amount"`
	Currency             string `json:"currency"`
}

/*
ChargesInformation entity.
*/
type ChargesInformation struct {
	ID                      uint            `json:"-"`
	AttributesID            uint            `json:"-"`
	BearerCode              string          `json:"bearer_code"`
	ReceiverChargesAmount   string          `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string          `json:"receiver_charges_currency"`
	SenderCharges           []SenderCharges `json:"sender_charges"`
}

/*
Fx entity.
*/
type Fx struct {
	ID                uint   `json:"-"`
	AttributesID      uint   `json:"-"`
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}
