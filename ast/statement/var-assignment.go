//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package statement

import "github.com/algotiqa/tiq-engine/ast/expression"

//=============================================================================
//===
//=== Variable declaration statement
//===
//=============================================================================

type VarAssignment struct {
	Chain      *expression.ChainedExpression
	Expression expression.Expression
}

//=============================================================================

func NewVarAssignment(chain *expression.ChainedExpression, expression expression.Expression) *VarAssignment {
	return &VarAssignment{
		Chain:      chain,
		Expression: expression,
	}
}

//=============================================================================

func (v *VarAssignment) Execute() error {
	return nil
}

//=============================================================================
