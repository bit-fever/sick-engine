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
	"github.com/algotiqa/tiq-engine/core/data"
)

//=============================================================================
//===
//=== If statement
//===
//=============================================================================

type IfStatement struct {
	condBlocks []*ConditionalBlock
}

//=============================================================================

func NewIfStatement() *IfStatement {
	return &IfStatement{}
}

//=============================================================================

func (s *IfStatement) AddConditionalBlock(cb *ConditionalBlock) {
	s.condBlocks = append(s.condBlocks, cb)
}

//=============================================================================

func (s *IfStatement) Execute() error {
	return nil
}

//=============================================================================
//===
//=== Conditional block
//===
//=============================================================================

type ConditionalBlock struct {
	condition expression.Expression
	block     *data.Block
}

//=============================================================================

func NewConditionalBlock(c expression.Expression, b *data.Block) *ConditionalBlock {
	return &ConditionalBlock{
		condition: c,
		block:     b,
	}
}

//=============================================================================
