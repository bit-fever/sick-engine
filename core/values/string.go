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
//=== String value
//===
//=============================================================================

type StringValue struct {
	value string
}

//=============================================================================

func NewStringValue(value string) *StringValue {
	return &StringValue{
		value: value,
	}
}

//=============================================================================

func (v *StringValue) Data() any {
	return v.value
}

//=============================================================================

func (v *StringValue) Type() types.Type {
	return types.NewStringType()
}

//=============================================================================

func (v *StringValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *StringValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *StringValue) String() string {
	return v.value
}

//=============================================================================
