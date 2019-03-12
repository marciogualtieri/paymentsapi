package repositories

import (
	"github.com/marciogualtieri/paymentsapi/models"
)

/*
PaymentsRepository defines an interface for the payments repository.
*/
//go:generate moq -out repository_test_moq_generated.go . PaymentsRepository
type PaymentsRepository interface {
	Create(payment models.Payment) (*string, error)
	Read(id string) (*models.Payment, error)
	Update(payment models.Payment) error
	Delete(id string) error
	ReadAll() ([]models.Payment, error)
	Close()
}
