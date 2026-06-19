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
	"github.com/algotiqa/tiq-engine/core/values"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Map expression
//===
//=============================================================================

type MapExpression struct {
	Values map[values.Value]Expression
	info   *parser.Info
}

//=============================================================================

func NewMapExpression(info *parser.Info) *MapExpression {
	return &MapExpression{
		Values: map[values.Value]Expression{},
		info:   info,
	}
}

//=============================================================================

func (e *MapExpression) Set(key values.Value, value Expression) {
	e.Values[key] = value
}

//=============================================================================

func (e *MapExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	return nil, errors.New("not implemented")
}

//=============================================================================

func (e *MapExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
