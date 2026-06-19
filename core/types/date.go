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
//=== Date
//===
//=============================================================================

type DateType struct {
}

//=============================================================================

var dateType = &DateType{}

//=============================================================================

func NewDateType() *DateType {
	return dateType
}

//=============================================================================

func (DateType) Code() int8 {
	return CodeDate
}

//=============================================================================

func (DateType) String() string {
	return "date"
}

//=============================================================================
