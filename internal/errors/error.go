package errors

type Type uint

const (
	Unknown  Type = 500
	Invalid  Type = 400
	NotFound Type = 404
)

type Error struct {
	Type    Type
	Message string
	Err     error
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Unwrap() error {
	return e.Err
}

var _ error = Error{}
