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

const (
	BarOpen  = 1
	BarHigh  = 2
	BarLow   = 3
	BarClose = 4
)

//=============================================================================
//===
//=== Bar Access
//===
//=============================================================================

type BarAccessExpression struct {
	Bar      int
	Accessor Expression
	info     *parser.Info
}

//=============================================================================

func NewBarAccessExpression(bar int, accessor Expression, info *parser.Info) *BarAccessExpression {
	return &BarAccessExpression{
		Bar:      bar,
		Accessor: accessor,
		info:     info,
	}
}

//=============================================================================

func (e *BarAccessExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t, err := e.Accessor.ResolveType(scope, embedder, depth)
	if err != nil {
		return nil, err
	}

	if t.Code() != types.CodeInt {
		return nil, errors.New("accessor's type must be int: '" + t.String() + "'")
	}

	return types.NewRealType(), nil
}

//=============================================================================

func (e *BarAccessExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
