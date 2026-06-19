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
//=== Type codes
//===
//=============================================================================

const (
	CodeInt        =  0
	CodeReal       =  1
	CodeBool       =  2
	CodeString     =  3
	CodeTime       =  4
	CodeDate       =  5
	CodeTimeseries =  6
	CodeEnum       =  7
	CodeClass      =  8
	CodeList       =  9
	CodeMap        = 10
	CodeError      = 11
	CodeToFindOut  = -1
)

//=============================================================================
//===
//=== Type interface
//===
//=============================================================================

type Type interface {
	Code()   int8
	String() string
}

//=============================================================================
