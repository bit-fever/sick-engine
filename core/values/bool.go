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
	"strconv"

	"github.com/algotiqa/tiq-engine/core/types"
)

//=============================================================================
//===
//=== Bool value
//===
//=============================================================================

type BoolValue struct {
	value bool
}

//=============================================================================

func NewBoolValue(value bool) *BoolValue {
	return &BoolValue{
		value: value,
	}
}

//=============================================================================

func (v *BoolValue) Data() any {
	return v.value
}

//=============================================================================

func (v *BoolValue) Type() types.Type {
	return types.NewBoolType()
}

//=============================================================================

func (v *BoolValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *BoolValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *BoolValue) String() string {
	return strconv.FormatBool(v.value)
}

//=============================================================================
