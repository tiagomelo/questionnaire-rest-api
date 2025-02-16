// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package validate

import (
	"encoding/json"
)

var jsonMarshal = json.Marshal

// FieldError is used to indicate an error with a specific request field.
type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// FieldErrors represents a collection of field errors.
type FieldErrors []FieldError

// Error implements the error interface.
func (fe FieldErrors) Error() string {
	d, err := jsonMarshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}
