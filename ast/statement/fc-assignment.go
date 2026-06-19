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
//=== Function call multiple assignment
//===
//=============================================================================

type FunctionCallAssignment struct {
	Chains       []*expression.ChainedExpression
	FunctionCall *expression.ChainedExpression
}

//=============================================================================

func NewFunctionCallAssignment(chains []*expression.ChainedExpression, fc *expression.ChainedExpression) *FunctionCallAssignment {
	return &FunctionCallAssignment{
		Chains:       chains,
		FunctionCall: fc,
	}
}

//=============================================================================

func (v *FunctionCallAssignment) Execute() error {
	return nil
}

//=============================================================================
