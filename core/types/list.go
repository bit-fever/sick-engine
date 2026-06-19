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
//=== List
//===
//=============================================================================

type ListType struct {
	SubType Type
}

//=============================================================================

func NewListType(t Type) *ListType {
	return &ListType{
		SubType: t,
	}
}

//=============================================================================

func (t *ListType) Code() int8 {
	return CodeList
}

//=============================================================================

func (t *ListType) String() string {
	return "list=("+ t.SubType.String() +")"
}

//=============================================================================
