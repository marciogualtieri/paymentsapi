package testing

import (
	. "github.com/marciogualtieri/paymentsapi/models"
	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func randomStringID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func randomNumberID() uint {
	return uint(rand.Uint64())
}

func randomizePaymentIDs(payment Payment) Payment {
	payment.ID = randomStringID()
	payment.Attributes.ID = randomNumberID()
	payment.Attributes.PaymentID = randomStringID()
	payment.Attributes.BeneficiaryParty.ID = randomNumberID()
	payment.Attributes.BeneficiaryParty.AttributesID = randomNumberID()
	payment.Attributes.DebtorParty.ID = randomNumberID()
	payment.Attributes.DebtorParty.AttributesID = randomNumberID()
	payment.Attributes.SponsorParty.ID = randomNumberID()
	payment.Attributes.SponsorParty.AttributesID = randomNumberID()
	payment.Attributes.ChargesInformation.ID = randomNumberID()
	payment.Attributes.ChargesInformation.AttributesID = randomNumberID()
	payment.Attributes.Fx.ID = randomNumberID()
	payment.Attributes.Fx.AttributesID = randomNumberID()
	for _, senderCharge := range payment.Attributes.ChargesInformation.SenderCharges {
		senderCharge.ID = randomNumberID()
		senderCharge.ChargesInformationID = randomNumberID()
	}
	return payment
}

func createMultipleRandomizedTestPayments(number int) []Payment {
	var payments []Payment
	for i := 0; i < number; i++ {
		payment := randomizePaymentIDs(TestPayment)
		payments = append(payments, payment)
	}
	return payments
}

func TestSpec(t *testing.T) {
	Convey("Given a string & a matching pattern.", t, func() {
		inputString := "123"
		inputPattern := "^[0-9]{3}$"
		Convey("When I match them using the regex matcher.", func() {
			result := ShouldMatchRegex(inputString, inputPattern)
			Convey("Then they do not match.", func() {
				So(result, ShouldBeEmpty)
			})
		})
	})

	Convey("Given a string & a not matching pattern.", t, func() {
		inputString := "123"
		inputPattern := "^[0-9]{2}$"
		Convey("When they are matched using the regex matcher.", func() {
			result := ShouldMatchRegex(inputString, inputPattern)
			Convey("Then they do not match.", func() {
				So(result, ShouldNotBeEmpty)
			})
		})
	})

	Convey("Given two resembling payment structs.", t, func() {
		payment := randomizePaymentIDs(TestPayment)
		anotherPayment := randomizePaymentIDs(TestPayment)
		Convey("When they are matched using the payment resemble matcher.", func() {
			result := ShouldPaymentResemble(payment, anotherPayment)
			Convey("Then payments match.", func() {
				So(result, ShouldBeEmpty)
			})
		})
	})

	Convey("Given two resembling payment structs.", t, func() {
		payment := randomizePaymentIDs(TestPayment)
		anotherPayment := randomizePaymentIDs(TestPayment)
		anotherPayment.Version = 123
		Convey("When they are matched using the payment resemble matcher.", func() {
			result := ShouldPaymentResemble(payment, anotherPayment)
			Convey("Then they do not match.", func() {
				So(result, ShouldNotBeEmpty)
			})
		})
	})

	Convey("Given two lists of resembling payment structs.", t, func() {
		aPaymentList := createMultipleRandomizedTestPayments(2)
		anotherPaymentList := createMultipleRandomizedTestPayments(2)

		Convey("When they are matched using the payment resemble matcher.", func() {
			result := ShouldPaymentsResemble(aPaymentList, anotherPaymentList)

			Convey("Then they match.", func() {
				So(result, ShouldBeEmpty)
			})
		})
	})

	Convey("Given two lists of payment structs of different lengths.", t, func() {
		aPaymentList := createMultipleRandomizedTestPayments(2)
		anotherPaymentList := createMultipleRandomizedTestPayments(3)

		Convey("When they are matched using the payment resemble matcher.", func() {
			result := ShouldPaymentsResemble(aPaymentList, anotherPaymentList)

			Convey("Then they do not match.", func() {
				So(result, ShouldNotBeEmpty)
			})
		})
	})

	Convey("Given two lists of not resembling payment structs.", t, func() {
		aPaymentList := createMultipleRandomizedTestPayments(2)
		aPaymentList[0].Type = "some-payment-type"
		anotherPaymentList := createMultipleRandomizedTestPayments(2)

		Convey("When they are matched using the payment resemble matcher.", func() {
			result := ShouldPaymentsResemble(aPaymentList, anotherPaymentList)

			Convey("Then they match.", func() {
				So(result, ShouldNotBeEmpty)
			})
		})
	})

}
