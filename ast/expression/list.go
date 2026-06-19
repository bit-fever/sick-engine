//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package expression

import (
	"errors"

	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== List expression
//===
//=============================================================================

type ListExpression struct {
	Values []Expression
	info   *parser.Info
}

//=============================================================================

func NewListExpression(info *parser.Info) *ListExpression {
	return &ListExpression{
		info: info,
	}
}

//=============================================================================

func (e *ListExpression) AddExpression(exp Expression) {
	e.Values = append(e.Values, exp)
}

//=============================================================================

func (e *ListExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	return nil, errors.New("not implemented")
}

//=============================================================================

func (e *ListExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
