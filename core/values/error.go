//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package values

import (
	"github.com/algotiqa/tiq-engine/core/types"
)

//=============================================================================
//===
//=== Error value
//===
//=============================================================================

type ErrorValue struct {
	value string
}

//=============================================================================

func NewErrorValue(value string) *ErrorValue {
	return &ErrorValue{
		value: value,
	}
}

//=============================================================================

func (v *ErrorValue) Data() any {
	return v.value
}

//=============================================================================

func (v *ErrorValue) Type() types.Type {
	return types.NewErrorType()
}

//=============================================================================

func (v *ErrorValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *ErrorValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *ErrorValue) String() string {
	return v.value
}

//=============================================================================
