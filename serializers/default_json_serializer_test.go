package serializers

import (
	. "github.com/marciogualtieri/paymentsapi/testing"
	"github.com/smartystreets/assertions"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultJSONSerializerSpec(t *testing.T) {

	Convey("Given a JSON serializer & a JSON payment.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentJSON := TestPaymentJSON

		Convey("When the payment JSON is parsed.", func() {
			paymentStruct, _ := jsonSerializer.ParsePayment(paymentJSON)

			Convey("The resulting payment struct is correct.", func() {
				So(*paymentStruct, ShouldResemble, TestPayment)
			})
		})
	})

	Convey("Given a JSON serializer & an invalid JSON payment.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentJSON := `{"not": "a", "valid": "payment"`

		Convey("When the invalid payment JSON is parsed.", func() {
			paymentStruct, err := jsonSerializer.ParsePayment(paymentJSON)

			Convey("The result is an error.", func() {
				So(paymentStruct, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a JSON serializer & a payment struct.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentStruct := TestPayment

		Convey("When the payment struct is serialized to a JSON.", func() {
			resultPaymentJSON := jsonSerializer.PaymentToString(paymentStruct)

			Convey("The resulting JSON payment is correct.", func() {
				So(resultPaymentJSON, assertions.ShouldEqualJSON, TestPaymentJSON)
			})
		})
	})

	Convey("Given a JSON serializer & a struct array of payments.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentStructArray, _ := jsonSerializer.ParsePayments(TestPaymentsArrayJSON)

		Convey("When the struct array of payments is serialized to JSON.", func() {
			jsonSerializer := NewDefaultJSONSerializer()
			resultPaymentsJSON := jsonSerializer.PaymentsToString(paymentStructArray)

			Convey("The resulting JSON array of payments is correct.", func() {
				So(resultPaymentsJSON, assertions.ShouldEqualJSON, TestPaymentsArrayJSON)
			})
		})
	})

	Convey("Given a JSON serializer & a JSON array of payments.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentsJSON := TestPaymentsArrayJSON

		Convey("When the JSON array of payments is parsed.", func() {
			paymentStructArray, _ := jsonSerializer.ParsePayments(paymentsJSON)

			Convey("The resulting struct array of payments is correct.", func() {
				resultPaymentsJSON := jsonSerializer.PaymentsToString(paymentStructArray)
				So(resultPaymentsJSON, assertions.ShouldEqualJSON, paymentsJSON)
			})
		})
	})

	Convey("Given a JSON serializer & an invalid JSON array of payments.", t, func() {
		jsonSerializer := NewDefaultJSONSerializer()
		paymentsJSON := `[{"not": "a"}, {"valid": "array"}, {"of": "payment"`

		Convey("When the invalid JSON array of payments is parsed.", func() {
			paymentStructArray, err := jsonSerializer.ParsePayments(paymentsJSON)

			Convey("The result is an error.", func() {
				So(paymentStructArray, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
