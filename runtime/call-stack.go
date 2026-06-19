//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package runtime

//=============================================================================
//===
//=== CallStack
//===
//=============================================================================

type CallStack struct {
	records []*CallRecord
	pointer int
}

//=============================================================================

func NewCallStack() *CallStack {
	cs := &CallStack{
		records: make([]*CallRecord, 1024),
	}

	cs.records[0] = NewCallRecord()
	return cs
}

//=============================================================================

func (cs *CallStack) ExistVariable(name string) bool {
	return cs.records[cs.pointer] != nil
}

//=============================================================================
