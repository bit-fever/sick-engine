//=============================================================================
//===
//=== Copyright (C) 2024-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package values

import (
	"github.com/algotiqa/tiq-engine/core/types"
	atypes "github.com/algotiqa/types"
)

//=============================================================================
//===
//=== Date value
//===
//=============================================================================

type DateValue struct {
	value atypes.Date
}

//=============================================================================

func NewDateValue(value atypes.Date) *DateValue {
	return &DateValue{
		value: value,
	}
}

//=============================================================================

func (v *DateValue) Data() any {
	return v.value
}

//=============================================================================

func (v *DateValue) Type() types.Type {
	return types.NewDateType()
}

//=============================================================================

func (v *DateValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *DateValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *DateValue) String() string {
	return v.value.String()
}

//=============================================================================
