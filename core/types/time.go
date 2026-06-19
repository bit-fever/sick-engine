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
//=== Time
//===
//=============================================================================

type TimeType struct {
}

//=============================================================================

var timeType = &TimeType{}

//=============================================================================

func NewTimeType() *TimeType {
	return timeType
}

//=============================================================================

func (TimeType) Code() int8 {
	return CodeTime
}

//=============================================================================

func (TimeType) String() string {
	return "time"
}

//=============================================================================
