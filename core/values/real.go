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
//=== Real value
//===
//=============================================================================

type RealValue struct {
	value float64
}

//=============================================================================

func NewRealValue(value float64) *RealValue {
	return &RealValue{
		value: value,
	}
}

//=============================================================================

func (v *RealValue) Data() any {
	return v.value
}

//=============================================================================

func (v *RealValue) Type() types.Type {
	return types.NewRealType()
}

//=============================================================================

func (v *RealValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *RealValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *RealValue) String() string {
	return strconv.FormatFloat(v.value, 'f', -1, 64)
}

//=============================================================================
