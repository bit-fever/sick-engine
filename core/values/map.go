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
//=== Map value
//===
//=============================================================================

type MapValue struct {
	values    map[Value]Value
	keyType   types.Type
	valueType types.Type
}

//=============================================================================

func NewMapValue(keyType, valueType types.Type) *MapValue {
	return &MapValue{
		keyType:   keyType,
		valueType: valueType,
		values:    map[Value]Value{},
	}
}

//=============================================================================

func (v *MapValue) Set(key Value, value Value) {
	v.values[key] = value
}

//=============================================================================

func (v *MapValue) Data() any {
	return v.values
}

//=============================================================================

func (v *MapValue) Type() types.Type {
	return v.keyType
}

//=============================================================================

func (v *MapValue) Equals(other Value) bool {
	return false
}

//=============================================================================

func (v *MapValue) LessThan(other Value) bool {
	return false
}

//=============================================================================

func (v *MapValue) String() string {
	var sb strings.Builder
	sb.WriteString("{ ")

	for key, value := range v.values {
		sb.WriteString(key.String())
		sb.WriteString(": ")
		sb.WriteString(value.String())
		sb.WriteString(", ")
	}

	sb.WriteString("}")
	return sb.String()
}

//=============================================================================
