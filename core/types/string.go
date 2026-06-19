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
//=== String
//===
//=============================================================================

type StringType struct {
}

//=============================================================================

var stringType = &StringType{}

//=============================================================================

func NewStringType() *StringType {
	return stringType
}

//=============================================================================

func (StringType) Code() int8 {
	return CodeString
}

//=============================================================================

func (StringType) String() string {
	return "string"
}

//=============================================================================
