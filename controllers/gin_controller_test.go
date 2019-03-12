package controllers

import (
	"errors"
	"fmt"
	. "github.com/marciogualtieri/paymentsapi/configuration"
	. "github.com/marciogualtieri/paymentsapi/errors"
	. "github.com/marciogualtieri/paymentsapi/models"
	. "github.com/marciogualtieri/paymentsapi/repositories"
	. "github.com/marciogualtieri/paymentsapi/serializers"
	. "github.com/marciogualtieri/paymentsapi/testing"
	"github.com/nu7hatch/gouuid"
	"github.com/smartystreets/assertions"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func pathWithID(id string) string {
	return fmt.Sprintf("/paymentsapi/%s", id)
}

func paymentsMapForMockup(payments []Payment) map[string]Payment {
	paymentMap := make(map[string]Payment)
	for _, payment := range payments {
		paymentMap[payment.ID] = payment
	}
	return paymentMap
}

func TestGinControllerSpec(t *testing.T) {
	serializer := NewDefaultJSONSerializer()
	configuration := GetTestConfiguration()
	paymentStructArray, _ := serializer.ParsePayments(TestPaymentsArrayJSON)
	paymentStructMap := paymentsMapForMockup(paymentStructArray)

	mockedRepository := &PaymentsRepositoryMock{
		ReadFunc: func(id string) (*Payment, error) {
			if payment, ok := paymentStructMap[id]; ok {
				return &payment, nil
			}
			return nil, NewErrRepositoryRecordNotFound()
		},
		CreateFunc: func(payment Payment) (*string, error) {
			id, _ := uuid.NewV4()
			payment.ID = id.String()
			paymentStructMap[payment.ID] = payment
			return &payment.ID, nil
		},
		UpdateFunc: func(payment Payment) error {
			if _, ok := paymentStructMap[payment.ID]; ok {
				paymentStructMap[payment.ID] = payment
				return nil
			}
			return NewErrRepositoryRecordNotFound()
		},
		DeleteFunc: func(id string) error {
			if payment, ok := paymentStructMap[id]; ok {
				delete(paymentStructMap, payment.ID)
				return nil
			}
			return NewErrRepositoryRecordNotFound()
		},
		ReadAllFunc: func() ([]Payment, error) {
			return paymentStructArray, nil
		},
	}

	mockedAlwaysFailsRepository := &PaymentsRepositoryMock{
		ReadFunc: func(id string) (*Payment, error) {
			return nil, errors.New("")
		},
		CreateFunc: func(payment Payment) (*string, error) {
			return nil, errors.New("")
		},
		UpdateFunc: func(payment Payment) error {
			return errors.New("")
		},
		DeleteFunc: func(id string) error {
			return errors.New("")
		},
		ReadAllFunc: func() ([]Payment, error) {
			return nil, errors.New("")
		},
	}

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When an existent payment is fetched.", func() {
			response := GetResource(engine, pathWithID(TestPayment.ID))

			Convey("Then the response has a payment & status <OK>.", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
				So(response.Body.String(), assertions.ShouldEqualJSON, TestPaymentJSON)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a non-existent payment is fetched.", func() {
			response := GetResource(engine, pathWithID(TestNonExistentPayment.ID))

			Convey("Then the response has status <NOT FOUND>.", func() {
				So(response.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})

	Convey("Given a Gin engine & a defective repository.", t, func() {
		engine := NewGinEngine(mockedAlwaysFailsRepository, serializer, configuration)

		Convey("When a payment is fetched.", func() {
			response := GetResource(engine, pathWithID(TestPayment.ID))

			Convey("Then the response has an error & status <INTERNAL SERVER ERROR>", func() {
				So(response.Code, ShouldEqual, http.StatusInternalServerError)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a payment is created.", func() {
			response := PostResource(engine, configuration.BaseResource, TestPaymentJSON)

			Convey("Then the response has an ID for the payment & status <CREATED>.", func() {
				So(response.Code, ShouldEqual, http.StatusCreated)
				So(response.Body.String(), ShouldMatchRegex, JSONIDRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a create for an invalid payment is attempted.", func() {
			response := PostResource(engine, configuration.BaseResource, TestInvalidPaymentJSON)

			Convey("Then the response has an error & status <BAD REQUEST>.", func() {
				So(response.Code, ShouldEqual, http.StatusBadRequest)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine & a defective repository.", t, func() {
		engine := NewGinEngine(mockedAlwaysFailsRepository, serializer, configuration)

		Convey("When a payment is created.", func() {
			response := PostResource(engine, configuration.BaseResource, TestPaymentJSON)

			Convey("Then the response has an error & status <INTERNAL SERVER ERROR>.", func() {
				So(response.Code, ShouldEqual, http.StatusInternalServerError)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a payment is updated.", func() {
			response := PutResource(engine, pathWithID(TestPayment.ID), TestPaymentJSON)

			Convey("The response status is <OK>.", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a payment that doesn't exist is updated.", func() {
			response := PutResource(engine, pathWithID(TestNonExistentPayment.ID), TestNonExistentPaymentJSON)

			Convey("Then the response has an error & status <NOT FOUND>", func() {
				So(response.Code, ShouldEqual, http.StatusNotFound)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When an invalid payment is updated.", func() {
			response := PutResource(engine, pathWithID(TestPayment.ID), TestInvalidPaymentJSON)

			Convey("Then the response has an error & status <BAD REQUEST>.", func() {
				So(response.Code, ShouldEqual, http.StatusBadRequest)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a payment is deleted.", func() {
			response := DeleteResource(engine, pathWithID(TestPayment.ID))

			Convey("The response status is <OK>.", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When a payment is deleted.", func() {
			response := DeleteResource(engine, pathWithID(TestNonExistentPayment.ID))

			Convey("Then the response has an error & status <NOT FOUND>", func() {
				So(response.Code, ShouldEqual, http.StatusNotFound)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

	Convey("Given a Gin engine.", t, func() {
		engine := NewGinEngine(mockedRepository, serializer, configuration)

		Convey("When payments are listed.", func() {
			response := GetResource(engine, configuration.BaseResource)

			Convey("Then the response has a payments & status <OK>.", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
				So(response.Body.String(), assertions.ShouldEqualJSON, TestPaymentsArrayJSON)
			})
		})
	})

	Convey("Given a Gin engine & a defective repository.", t, func() {
		engine := NewGinEngine(mockedAlwaysFailsRepository, serializer, configuration)

		Convey("When payments are listed.", func() {
			response := GetResource(engine, configuration.BaseResource)

			Convey("Then the response has an error & status <INTERNAL SERVER ERROR>", func() {
				So(response.Code, ShouldEqual, http.StatusInternalServerError)
				So(response.Body.String(), ShouldMatchRegex, JSONErrorRegex)
			})
		})
	})

}
