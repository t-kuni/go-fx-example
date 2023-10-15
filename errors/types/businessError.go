package types

type BasicBusinessError struct {
	Message string
	Params  map[string]interface{}
}

func NewBasicBusinessError(message string, params map[string]interface{}) error {
	return BasicBusinessError{
		Message: message,
		Params:  params,
	}
}

func (e BasicBusinessError) Error() string {
	return e.Message
}
