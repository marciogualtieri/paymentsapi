// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repositories

import (
	"github.com/marciogualtieri/paymentsapi/models"
	"sync"
)

var (
	lockPaymentsRepositoryMockClose   sync.RWMutex
	lockPaymentsRepositoryMockCreate  sync.RWMutex
	lockPaymentsRepositoryMockDelete  sync.RWMutex
	lockPaymentsRepositoryMockRead    sync.RWMutex
	lockPaymentsRepositoryMockReadAll sync.RWMutex
	lockPaymentsRepositoryMockUpdate  sync.RWMutex
)

// Ensure, that PaymentsRepositoryMock does implement PaymentsRepository.
// If this is not the case, regenerate this file with moq.
var _ PaymentsRepository = &PaymentsRepositoryMock{}

// PaymentsRepositoryMock is a mock implementation of PaymentsRepository.
//
//     func TestSomethingThatUsesPaymentsRepository(t *testing.T) {
//
//         // make and configure a mocked PaymentsRepository
//         mockedPaymentsRepository := &PaymentsRepositoryMock{
//             CloseFunc: func()  {
// 	               panic("mock out the Close method")
//             },
//             CreateFunc: func(payment models.Payment) (*string, error) {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(id string) error {
// 	               panic("mock out the Delete method")
//             },
//             ReadFunc: func(id string) (*models.Payment, error) {
// 	               panic("mock out the Read method")
//             },
//             ReadAllFunc: func() ([]models.Payment, error) {
// 	               panic("mock out the ReadAll method")
//             },
//             UpdateFunc: func(payment models.Payment) error {
// 	               panic("mock out the Update method")
//             },
//         }
//
//         // use mockedPaymentsRepository in code that requires PaymentsRepository
//         // and then make assertions.
//
//     }
type PaymentsRepositoryMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func()

	// CreateFunc mocks the Create method.
	CreateFunc func(payment models.Payment) (*string, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(id string) error

	// ReadFunc mocks the Read method.
	ReadFunc func(id string) (*models.Payment, error)

	// ReadAllFunc mocks the ReadAll method.
	ReadAllFunc func() ([]models.Payment, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(payment models.Payment) error

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Payment is the payment argument value.
			Payment models.Payment
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// ID is the id argument value.
			ID string
		}
		// Read holds details about calls to the Read method.
		Read []struct {
			// ID is the id argument value.
			ID string
		}
		// ReadAll holds details about calls to the ReadAll method.
		ReadAll []struct {
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Payment is the payment argument value.
			Payment models.Payment
		}
	}
}

// Close calls CloseFunc.
func (mock *PaymentsRepositoryMock) Close() {
	if mock.CloseFunc == nil {
		panic("PaymentsRepositoryMock.CloseFunc: method is nil but PaymentsRepository.Close was just called")
	}
	callInfo := struct {
	}{}
	lockPaymentsRepositoryMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockPaymentsRepositoryMockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedPaymentsRepository.CloseCalls())
func (mock *PaymentsRepositoryMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockPaymentsRepositoryMockClose.RLock()
	calls = mock.calls.Close
	lockPaymentsRepositoryMockClose.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *PaymentsRepositoryMock) Create(payment models.Payment) (*string, error) {
	if mock.CreateFunc == nil {
		panic("PaymentsRepositoryMock.CreateFunc: method is nil but PaymentsRepository.Create was just called")
	}
	callInfo := struct {
		Payment models.Payment
	}{
		Payment: payment,
	}
	lockPaymentsRepositoryMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockPaymentsRepositoryMockCreate.Unlock()
	return mock.CreateFunc(payment)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedPaymentsRepository.CreateCalls())
func (mock *PaymentsRepositoryMock) CreateCalls() []struct {
	Payment models.Payment
} {
	var calls []struct {
		Payment models.Payment
	}
	lockPaymentsRepositoryMockCreate.RLock()
	calls = mock.calls.Create
	lockPaymentsRepositoryMockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *PaymentsRepositoryMock) Delete(id string) error {
	if mock.DeleteFunc == nil {
		panic("PaymentsRepositoryMock.DeleteFunc: method is nil but PaymentsRepository.Delete was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockPaymentsRepositoryMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockPaymentsRepositoryMockDelete.Unlock()
	return mock.DeleteFunc(id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedPaymentsRepository.DeleteCalls())
func (mock *PaymentsRepositoryMock) DeleteCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockPaymentsRepositoryMockDelete.RLock()
	calls = mock.calls.Delete
	lockPaymentsRepositoryMockDelete.RUnlock()
	return calls
}

// Read calls ReadFunc.
func (mock *PaymentsRepositoryMock) Read(id string) (*models.Payment, error) {
	if mock.ReadFunc == nil {
		panic("PaymentsRepositoryMock.ReadFunc: method is nil but PaymentsRepository.Read was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockPaymentsRepositoryMockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	lockPaymentsRepositoryMockRead.Unlock()
	return mock.ReadFunc(id)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedPaymentsRepository.ReadCalls())
func (mock *PaymentsRepositoryMock) ReadCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockPaymentsRepositoryMockRead.RLock()
	calls = mock.calls.Read
	lockPaymentsRepositoryMockRead.RUnlock()
	return calls
}

// ReadAll calls ReadAllFunc.
func (mock *PaymentsRepositoryMock) ReadAll() ([]models.Payment, error) {
	if mock.ReadAllFunc == nil {
		panic("PaymentsRepositoryMock.ReadAllFunc: method is nil but PaymentsRepository.ReadAll was just called")
	}
	callInfo := struct {
	}{}
	lockPaymentsRepositoryMockReadAll.Lock()
	mock.calls.ReadAll = append(mock.calls.ReadAll, callInfo)
	lockPaymentsRepositoryMockReadAll.Unlock()
	return mock.ReadAllFunc()
}

// ReadAllCalls gets all the calls that were made to ReadAll.
// Check the length with:
//     len(mockedPaymentsRepository.ReadAllCalls())
func (mock *PaymentsRepositoryMock) ReadAllCalls() []struct {
} {
	var calls []struct {
	}
	lockPaymentsRepositoryMockReadAll.RLock()
	calls = mock.calls.ReadAll
	lockPaymentsRepositoryMockReadAll.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *PaymentsRepositoryMock) Update(payment models.Payment) error {
	if mock.UpdateFunc == nil {
		panic("PaymentsRepositoryMock.UpdateFunc: method is nil but PaymentsRepository.Update was just called")
	}
	callInfo := struct {
		Payment models.Payment
	}{
		Payment: payment,
	}
	lockPaymentsRepositoryMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockPaymentsRepositoryMockUpdate.Unlock()
	return mock.UpdateFunc(payment)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedPaymentsRepository.UpdateCalls())
func (mock *PaymentsRepositoryMock) UpdateCalls() []struct {
	Payment models.Payment
} {
	var calls []struct {
		Payment models.Payment
	}
	lockPaymentsRepositoryMockUpdate.RLock()
	calls = mock.calls.Update
	lockPaymentsRepositoryMockUpdate.RUnlock()
	return calls
}
