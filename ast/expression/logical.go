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
//=== And
//===
//=============================================================================

type AndExpression struct {
	expressions []Expression
	info        *parser.Info
}

//=============================================================================

func NewAndExpression(expressions []Expression, info *parser.Info) *AndExpression {
	return &AndExpression{
		expressions: expressions,
		info:        info,
	}
}

//=============================================================================

func (e *AndExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	for _, ex := range e.expressions {
		err := checkBoolean(ex, scope, embedder, depth)
		if err != nil {
			return nil, err
		}
	}

	return types.NewBoolType(), nil
}

//=============================================================================

func (e *AndExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Or
//===
//=============================================================================

type OrExpression struct {
	expressions []Expression
	info        *parser.Info
}

//=============================================================================

func NewOrExpression(expressions []Expression, info *parser.Info) *OrExpression {
	return &OrExpression{
		expressions: expressions,
		info:        info,
	}
}

//=============================================================================

func (e *OrExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	for _, ex := range e.expressions {
		err := checkBoolean(ex, scope, embedder, depth)
		if err != nil {
			return nil, err
		}
	}

	return types.NewBoolType(), nil
}

//=============================================================================

func (e *OrExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Not
//===
//=============================================================================

type NotExpression struct {
	expression Expression
	info       *parser.Info
}

//=============================================================================

func NewNotExpression(e Expression, info *parser.Info) *NotExpression {
	return &NotExpression{
		expression: e,
		info:       info,
	}
}

//=============================================================================

func (e *NotExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	err := checkBoolean(e.expression, scope, embedder, depth)
	if err != nil {
		return nil, err
	}

	return types.NewBoolType(), nil
}

//=============================================================================

func (e *NotExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func checkBoolean(e Expression, scope interfaces.Scope, embedder interfaces.Symbol, depth int) error {
	t, err := e.ResolveType(scope, embedder, depth)
	if err != nil {
		return err
	}

	if t.Code() != types.CodeBool {
		return errors.New("resulting expression is not boolean")
	}

	return nil
}

//=============================================================================
