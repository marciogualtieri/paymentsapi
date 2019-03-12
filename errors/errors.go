package errors

/*
ErrParsingJSON is a custom error for when JSON can' t be parsed.
*/
type ErrParsingJSON struct {
	message string
}

/*
NewErrParsingJSON creates a new ErrItemNotFound error.
*/
func NewErrParsingJSON(message string) error {
	return &ErrParsingJSON{
		message: message,
	}
}

func (e *ErrParsingJSON) Error() string {
	return e.message
}

/*
ErrRepository is a custom error for repository errors.
*/
type ErrRepository struct {
	message string
}

/*
NewErrRepository creates a new ErrItemNotFound error.
*/
func NewErrRepository(message string) error {
	return &ErrRepository{
		message: message,
	}
}

func (e *ErrRepository) Error() string {
	return e.message
}

/*
ErrRepositoryRecordNotFound is a custom error for repository errors.
*/
type ErrRepositoryRecordNotFound struct {
	message string
}

/*
NewErrRepositoryRecordNotFound creates a new ErrItemNotFound error.
*/
func NewErrRepositoryRecordNotFound() error {
	return &ErrRepositoryRecordNotFound{
		message: "record not found",
	}
}

func (e *ErrRepositoryRecordNotFound) Error() string {
	return e.message
}
