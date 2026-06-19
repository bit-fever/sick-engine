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
//=== Time value
//===
//=============================================================================

type TimeValue struct {
	value atypes.Time
}

//=============================================================================

func NewTimeValue(value atypes.Time) *TimeValue {
	return &TimeValue{
		value: value,
	}
}

//=============================================================================

func (v *TimeValue) Data() any {
	return v.value
}

//=============================================================================

func (v *TimeValue) Type() types.Type {
	return types.NewTimeType()
}

//=============================================================================

func (v *TimeValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *TimeValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *TimeValue) String() string {
	return v.value.String()
}

//=============================================================================
