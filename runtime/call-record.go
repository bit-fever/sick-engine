//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package runtime

import (
	"github.com/algotiqa/tiq-engine/core/values"
)

//=============================================================================
//===
//=== CallRecord
//===
//=============================================================================

type CallRecord struct {
	vars map[string]values.Value
}

//=============================================================================

func NewCallRecord() *CallRecord {
	return &CallRecord{
		vars: make(map[string]values.Value),
	}
}

//=============================================================================

func (cr *CallRecord) ExistsVariable(name string) bool {
	return false
}

//=============================================================================

func (cr *CallRecord) SetVariable(name, string, value values.Value) {

}

//=============================================================================
