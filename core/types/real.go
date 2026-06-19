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
//=== Real
//===
//=============================================================================

type RealType struct {
}

//=============================================================================

var realType = &RealType{}

//=============================================================================

func NewRealType() *RealType {
	return realType
}

//=============================================================================

func (RealType) Code() int8 {
	return CodeReal
}

//=============================================================================

func (RealType) String() string {
	return "real"
}

//=============================================================================
