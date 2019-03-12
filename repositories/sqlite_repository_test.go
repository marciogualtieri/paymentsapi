package repositories

import (
	. "github.com/marciogualtieri/paymentsapi/configuration"
	. "github.com/marciogualtieri/paymentsapi/errors"
	. "github.com/marciogualtieri/paymentsapi/models"
	. "github.com/marciogualtieri/paymentsapi/testing"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func createMultipleTestPayments(repository *SqlitePaymentRepository, number int) []Payment {
	var payments []Payment
	for i := 0; i < number; i++ {
		id, _ := repository.Create(TestPayment)
		payment, _ := repository.Read(*id)
		payments = append(payments, *payment)
	}
	return payments
}

func TestSqliteRepositorySpec(t *testing.T) {

	Convey("Given a Sqlite repository & an existent payment.", t, func() {
		repository := NewSqlitePaymentRepository(GetTestConfiguration())
		id, _ := repository.Create(TestPayment)

		Convey("When a payment is read.", func() {
			payment, err := repository.Read(*id)

			Convey("Then the result is the payment.", func() {
				So(err, ShouldBeNil)
				So(*payment, ShouldPaymentResemble, TestPayment)
			})
		})
	})

	Convey("Given a Sqlite repository & an existent payment.", t, func() {
		repository := NewSqlitePaymentRepository(GetTestConfiguration())
		id, _ := repository.Create(TestPayment)
		payment, _ := repository.Read(*id)

		Convey("When a payment is updated.", func() {
			payment.OrganisationID = "changed-with-some-organization-id"
			err := repository.Update(*payment)

			Convey("Then the result payment is up-to-date.", func() {
				So(err, ShouldBeNil)
				updatedPayment, _ := repository.Read(*id)
				So(*updatedPayment, ShouldPaymentResemble, *payment)
			})
		})
	})

	Convey("Given a Sqlite repository & an existent payment.", t, func() {
		repository := NewSqlitePaymentRepository(GetTestConfiguration())
		id, _ := repository.Create(TestPayment)

		Convey("When a payment is deleted.", func() {
			err := repository.Delete(*id)

			Convey("Then the payment no longer exists.", func() {
				So(err, ShouldBeNil)
				_, err := repository.Read(*id)
				So(err, ShouldBeError, NewErrRepositoryRecordNotFound())
			})
		})
	})

	Convey("Given a Sqlite repository & an inexistent payment.", t, func() {
		repository := NewSqlitePaymentRepository(GetTestConfiguration())
		id := TestNonExistentPayment.ID

		Convey("When a payment is deleted.", func() {
			err := repository.Delete(id)

			Convey("Then the result is a record not found error.", func() {
				So(err, ShouldBeError, NewErrRepositoryRecordNotFound())
			})
		})
	})

	Convey("Given a Sqlite repository & a list of existent payments.", t, func() {
		paymentsSize := 3
		repository := NewSqlitePaymentRepository(GetTestConfiguration())
		expectedPayments := createMultipleTestPayments(repository, paymentsSize)

		Convey("When payments are listed.", func() {
			payments, err := repository.ReadAll()

			Convey("Then the result is the list of payments.", func() {
				So(err, ShouldBeNil)
				So(payments, ShouldHaveLength, paymentsSize)
				So(payments, ShouldPaymentsResemble, expectedPayments)
			})
		})
	})
}
