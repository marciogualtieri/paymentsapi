package errors

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func switchByType(err error) string {
	result := "not switched"
	switch err.(type) {
	case *ErrRepository:
		result = "switched to ErrRepository"
	case *ErrParsingJSON:
		result = "switched to ErrParsingJSON"
	case *ErrRepositoryRecordNotFound:
		result = "switched to ErrRepositoryRecordNotFound"
	}
	return result
}

func TestCustomErrorsSpec(t *testing.T) {

	Convey("Given a ErrRepository error.", t, func() {
		err := NewErrRepository("some repository error")

		Convey("When the error is switched by type.", func() {
			result := switchByType(err)

			Convey("The the error is switched correctly.", func() {
				So(result, ShouldEqual, "switched to ErrRepository")
			})
		})
	})

	Convey("Given a ErrParsingJSON error.", t, func() {
		err := NewErrParsingJSON("some parsing error")

		Convey("When the error is switched by type.", func() {
			result := switchByType(err)

			Convey("The the error is switched correctly.", func() {
				So(result, ShouldEqual, "switched to ErrParsingJSON")
			})
		})
	})

	Convey("Given a ErrRepositoryRecordNotFound error.", t, func() {
		err := NewErrRepositoryRecordNotFound()

		Convey("When the error is switched by type.", func() {
			result := switchByType(err)

			Convey("The the error is switched correctly.", func() {
				So(result, ShouldEqual, "switched to ErrRepositoryRecordNotFound")
			})
		})
	})
}
