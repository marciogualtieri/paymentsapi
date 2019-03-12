package testing

import (
	"github.com/marciogualtieri/paymentsapi/models"
	"io/ioutil"
)

func fileAsString(resourcePath string) string {
	testPaymentBytes, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		panic(err)
	}
	return string(testPaymentBytes)
}

/*
TestPaymentJSON is a single existent JSON payment.
*/
var TestPaymentJSON = fileAsString("../testing/resources/test_payment.json")

/*
TestPaymentsArrayJSON contains a JSON array of payments.
*/
var TestPaymentsArrayJSON = fileAsString("../testing/resources/test_payments_array.json")

/*
TestNonExistentPaymentJSON is a non-existent payment.
*/
var TestNonExistentPaymentJSON = fileAsString("../testing/resources/test_non_existent_payment.json")

/*
TestPayment is a payment struct.
*/
var TestPayment = models.Payment{
	Type:           "Payment",
	ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
	Version:        0,
	OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	Attributes: models.Attributes{
		Amount: "100.21",
		BeneficiaryParty: models.BeneficiaryParty{
			AccountName:       "W Owens",
			AccountNumber:     "31926819",
			AccountNumberCode: "BBAN",
			AccountType:       0,
			Address:           "1 The Beneficiary Localtown SE2",
			BankID:            "403000",
			BankIDCode:        "GBDSC",
			Name:              "Wilfred Jeremiah Owens",
		},
		ChargesInformation: models.ChargesInformation{
			BearerCode: "SHAR",
			SenderCharges: []models.SenderCharges{
				models.SenderCharges{
					Amount:   "5.00",
					Currency: "GBP",
				},
				models.SenderCharges{
					Amount:   "10.00",
					Currency: "USD",
				},
			},
			ReceiverChargesAmount:   "1.00",
			ReceiverChargesCurrency: "USD",
		},
		Currency: "GBP",
		DebtorParty: models.DebtorParty{
			AccountName:       "EJ Brown Black",
			AccountNumber:     "GB29XABC10161234567801",
			AccountNumberCode: "IBAN",
			Address:           "10 Debtor Crescent Sourcetown NE1",
			BankID:            "203301",
			BankIDCode:        "GBDSC",
			Name:              "Emelia Jane Brown",
		},
		EndToEndReference: "Wil piano Jan",
		Fx: models.Fx{
			ContractReference: "FX123",
			ExchangeRate:      "2.00000",
			OriginalAmount:    "200.42",
			OriginalCurrency:  "USD",
		},
		NumericReference:     "1002001",
		PaymentIdentifier:    "123456789012345678",
		PaymentPurpose:       "Paying for goods/services",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Payment for Em's piano lessons",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",
		SponsorParty: models.SponsorParty{
			AccountNumber: "56781234",
			BankID:        "123123",
			BankIDCode:    "GBDSC",
		},
	},
}

/*
TestNonExistentPayment is a non-existent payment struct.
*/
var TestNonExistentPayment = models.Payment{
	Type:           "Payment",
	ID:             "non-existent-payment-id",
	Version:        0,
	OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	Attributes: models.Attributes{
		Amount: "100.21",
		BeneficiaryParty: models.BeneficiaryParty{
			AccountName:       "W Owens",
			AccountNumber:     "31926819",
			AccountNumberCode: "BBAN",
			AccountType:       0,
			Address:           "1 The Beneficiary Localtown SE2",
			BankID:            "403000",
			BankIDCode:        "GBDSC",
			Name:              "Wilfred Jeremiah Owens",
		},
		ChargesInformation: models.ChargesInformation{
			BearerCode: "SHAR",
			SenderCharges: []models.SenderCharges{
				models.SenderCharges{
					Amount:   "5.00",
					Currency: "GBP",
				},
				models.SenderCharges{
					Amount:   "10.00",
					Currency: "USD",
				},
			},
			ReceiverChargesAmount:   "1.00",
			ReceiverChargesCurrency: "USD",
		},
		Currency: "GBP",
		DebtorParty: models.DebtorParty{
			AccountName:       "EJ Brown Black",
			AccountNumber:     "GB29XABC10161234567801",
			AccountNumberCode: "IBAN",
			Address:           "10 Debtor Crescent Sourcetown NE1",
			BankID:            "203301",
			BankIDCode:        "GBDSC",
			Name:              "Emelia Jane Brown",
		},
		EndToEndReference: "Wil piano Jan",
		Fx: models.Fx{
			ContractReference: "FX123",
			ExchangeRate:      "2.00000",
			OriginalAmount:    "200.42",
			OriginalCurrency:  "USD",
		},
		NumericReference:     "1002001",
		PaymentIdentifier:    "123456789012345678",
		PaymentPurpose:       "Paying for goods/services",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Payment for Em's piano lessons",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",
		SponsorParty: models.SponsorParty{
			AccountNumber: "56781234",
			BankID:        "123123",
			BankIDCode:    "GBDSC",
		},
	},
}

/*
UUIDRegex is the regex for UUID.
*/
var UUIDRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"

/*
JSONIDRegex is the regex for the JSON response for the payment ID.
*/
var JSONIDRegex = `\{\"ID\"\:\"` + UUIDRegex + `\"\}`

/*
JSONErrorRegex is the regex for the JSON response for an error.
*/
var JSONErrorRegex = `\{\"error\"\:\".*\"\}`

/*
TestInvalidPaymentJSON is invalid JSON.
*/
var TestInvalidPaymentJSON = "not-valid-paymento-json"
