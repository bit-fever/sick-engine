//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package statement

import (
	"github.com/algotiqa/tiq-engine/ast/expression"
	"github.com/algotiqa/tiq-engine/core/types"
)

//=============================================================================
//===
//=== Variable declaration statement
//===
//=============================================================================

type VarDeclaration struct {
	Name       string
	Type       types.Type
	Expression expression.Expression
}

//=============================================================================

func NewVarDeclaration(name string, type_ types.Type, e expression.Expression) *VarDeclaration {
	return &VarDeclaration{
		Name:       name,
		Type:       type_,
		Expression: e,
	}
}

//=============================================================================

func (v *VarDeclaration) Execute() error {
	return nil
}

//=============================================================================
