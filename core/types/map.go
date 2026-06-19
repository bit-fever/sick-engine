//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

//=============================================================================
//===
//=== Map
//===
//=============================================================================

type MapType struct {
	KeyType   Type
	ValueType Type
}

//=============================================================================

func NewMapType(keyType, valueType Type) *MapType {
	return &MapType{
		KeyType  : keyType,
		ValueType: valueType,
	}
}

//=============================================================================

func (t *MapType) Code() int8 {
	return CodeMap
}

//=============================================================================

func (t *MapType) String() string {
	return "map=("+ t.KeyType.String() +":"+ t.ValueType.String() +")"
}

//=============================================================================
