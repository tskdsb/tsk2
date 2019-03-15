package util

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type ErrorType string
type ErrorTemplate string

type TypeError struct {
	Err      error
	Type     ErrorType
	Template ErrorTemplate
	Message  string
	File     string
	Line     int
}

func (te *TypeError) Error() string {
	errData, _ := json.Marshal(te)
	return fmt.Sprintf("%s", errData)
}

func NewTypeError(err error) TypeError {
	var te = TypeError{
		Err: err,
	}

	_, te.File, te.Line, _ = runtime.Caller(2)

	return te
}
