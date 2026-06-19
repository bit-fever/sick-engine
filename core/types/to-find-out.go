//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

import (
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== ToFindOut
//===
//=============================================================================

type ToFindOutType struct {
	Name *data.FQIdentifier
	Info *parser.Info
}

//=============================================================================

func NewToFindOutType(name *data.FQIdentifier, info *parser.Info) *ToFindOutType {
	return &ToFindOutType{
		Name: name,
		Info: info,
	}
}

//=============================================================================

func (t *ToFindOutType) Code() int8 {
	return CodeToFindOut
}

//=============================================================================

func (t *ToFindOutType) String() string {
	return "tfo=(" + t.Name.String() + ")"
}

//=============================================================================
