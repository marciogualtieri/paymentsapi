package serializers

import (
	"github.com/marciogualtieri/paymentsapi/models"
)

/*
JSONSerializer defines an interface for serialization.
*/
type JSONSerializer interface {
	ParsePayment(paymentJSON string) (*models.Payment, error)
	ParsePayments(paymentJSON string) ([]models.Payment, error)
	PaymentsToString(payments []models.Payment) string
	PaymentToString(payment models.Payment) string
}
