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
//=== Int value
//===
//=============================================================================

type IntValue struct {
	value int64
}

//=============================================================================

func NewIntValue(value int64) *IntValue {
	return &IntValue{
		value: value,
	}
}

//=============================================================================

func (v *IntValue) Data() any {
	return v.value
}

//=============================================================================

func (v *IntValue) Type() types.Type {
	return types.NewIntType()
}

//=============================================================================

func (v *IntValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *IntValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *IntValue) String() string {
	return strconv.FormatInt(v.value, 10)
}

//=============================================================================
