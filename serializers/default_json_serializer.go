package serializers

import (
	"encoding/json"
	"github.com/marciogualtieri/paymentsapi/errors"
	"github.com/marciogualtieri/paymentsapi/models"
)

/*
DefaultJSONSerializer is the default implementation of the json serializer interface.
*/
type DefaultJSONSerializer struct{}

/*
NewDefaultJSONSerializer creates a new JSON serializer.
*/
func NewDefaultJSONSerializer() *DefaultJSONSerializer {
	return &DefaultJSONSerializer{}
}

/*
ParsePayment parses a JSON payments string.
*/
func (*DefaultJSONSerializer) ParsePayment(paymentJSON string) (*models.Payment, error) {
	var result models.Payment
	err := json.Unmarshal([]byte(paymentJSON), &result)
	if err != nil {
		return nil, errors.NewErrParsingJSON(err.Error())
	}
	return &result, nil
}

/*
ParsePayments parses a JSON payments string.
*/
func (*DefaultJSONSerializer) ParsePayments(paymentJSON string) ([]models.Payment, error) {
	var result []models.Payment
	err := json.Unmarshal([]byte(paymentJSON), &result)
	if err != nil {
		return nil, errors.NewErrParsingJSON(err.Error())
	}
	return result, nil
}

/*
PaymentsToString converts a payment array to a JSON string.
*/
func (*DefaultJSONSerializer) PaymentsToString(payments []models.Payment) string {
	byteResult, _ := json.Marshal(payments)
	return string(byteResult)
}

/*
PaymentToString converts a payment array to a JSON string.
*/
func (*DefaultJSONSerializer) PaymentToString(payment models.Payment) string {
	byteResult, _ := json.Marshal(payment)
	return string(byteResult)
}
