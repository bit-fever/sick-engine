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
	"strings"

	"github.com/algotiqa/tiq-engine/core/types"
)

//=============================================================================
//===
//=== List value
//===
//=============================================================================

type ListValue struct {
	values []Value
	type_  types.Type
}

//=============================================================================

func NewListValue(type_ types.Type) *ListValue {
	return &ListValue{
		type_: type_,
	}
}

//=============================================================================

func (v *ListValue) AddValue(value Value) {
	v.values = append(v.values, value)
}

//=============================================================================

func (v *ListValue) Data() any {
	return v.values
}

//=============================================================================

func (v *ListValue) Type() types.Type {
	return v.type_
}

//=============================================================================

func (v *ListValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *ListValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *ListValue) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for i, value := range v.values {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(value.String())
	}

	sb.WriteString("]")
	return sb.String()
}

//=============================================================================
