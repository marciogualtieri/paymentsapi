package testing

import (
	"fmt"
	"github.com/marciogualtieri/paymentsapi/models"
	"github.com/smartystreets/goconvey/convey"
	"regexp"
)

// TODO: Refactor ShouldPaymentResemble & ShouldPaymentResemble so the assertion generalizes to any struct using reflection.

/*
ShouldMatchRegex is a custom assertion for regex.
*/
func ShouldMatchRegex(actual interface{}, expected ...interface{}) string {
	value := actual.(string)
	pattern := expected[0].(string)
	ok, err := regexp.MatchString(pattern, value)
	if err == nil && ok {
		return ""
	}
	return fmt.Sprintf("%s should match expression %s", value, pattern)
}

func zeroPaymentIDs(payment models.Payment) models.Payment {
	payment.ID = ""
	payment.Attributes.ID = 0
	payment.Attributes.PaymentID = ""
	payment.Attributes.BeneficiaryParty.ID = 0
	payment.Attributes.BeneficiaryParty.AttributesID = 0
	payment.Attributes.DebtorParty.ID = 0
	payment.Attributes.DebtorParty.AttributesID = 0
	payment.Attributes.SponsorParty.ID = 0
	payment.Attributes.SponsorParty.AttributesID = 0
	payment.Attributes.ChargesInformation.ID = 0
	payment.Attributes.ChargesInformation.AttributesID = 0
	payment.Attributes.Fx.ID = 0
	payment.Attributes.Fx.AttributesID = 0
	for i := range payment.Attributes.ChargesInformation.SenderCharges {
		payment.Attributes.ChargesInformation.SenderCharges[i].ID = 0
		payment.Attributes.ChargesInformation.SenderCharges[i].ChargesInformationID = 0
	}
	return payment
}

/*
ShouldPaymentResemble is a custom assertion that ignores ID fields.
*/
func ShouldPaymentResemble(actual interface{}, expected ...interface{}) string {
	actualPayment := actual.(models.Payment)
	expectedPayment := expected[0].(models.Payment)
	actualPayment = zeroPaymentIDs(actualPayment)
	expectedPayment = zeroPaymentIDs(expectedPayment)
	return convey.ShouldResemble(actualPayment, expectedPayment)
}

/*
ShouldPaymentsResemble is a custom assertion that ignores ID fields.
*/
func ShouldPaymentsResemble(actual interface{}, expected ...interface{}) string {
	actualPayments := actual.([]models.Payment)
	expectedPayments := expected[0].([]models.Payment)
	if len(actualPayments) != len(expectedPayments) {
		return fmt.Sprintf("should be the same length:\nActual %+v\n Expected: %+v",
			actualPayments, expectedPayments)
	}
	result := ""
	for i := range actualPayments {
		result += ShouldPaymentResemble(actualPayments[i], expectedPayments[i])
	}
	return result
}
