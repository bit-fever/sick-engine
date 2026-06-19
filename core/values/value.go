//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package values

import "github.com/algotiqa/tiq-engine/core/types"

//=============================================================================
//===
//=== Value interface
//===
//=============================================================================

type Value interface {
	Data() any
	Type() types.Type

	Equals(other Value) bool
	LessThan(other Value) bool
	String() string
}

//=============================================================================
