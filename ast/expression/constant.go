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
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/core/values"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== ConstantValue
//===
//=============================================================================

type ConstantValueExpression struct {
	Value values.Value
	info  *parser.Info
}

//=============================================================================

func NewConstantValueExpression(value values.Value, info *parser.Info) *ConstantValueExpression {
	return &ConstantValueExpression{
		Value: value,
		info:  info,
	}
}

//=============================================================================

func (e *ConstantValueExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	return e.Value.Type(), nil
}

//=============================================================================

func (e *ConstantValueExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
