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
//=== Return statement
//===
//=============================================================================

type ReturnStatement struct {
	values []expression.Expression
}

//=============================================================================

func NewReturnStatement() *ReturnStatement {
	return &ReturnStatement{}
}

//=============================================================================

func (s *ReturnStatement) AddExpression(e expression.Expression) {
	s.values = append(s.values, e)
}

//=============================================================================

func (s *ReturnStatement) Execute() error {
	return nil
}

//=============================================================================
