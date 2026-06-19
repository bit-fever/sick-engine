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
//=== Operands
//===
//=============================================================================

const (
	RelOpEqual          = "="
	RelOpNotEqual       = "<>"
	RelOpLessThan       = "<"
	RelOpGreaterThan    = ">"
	RelOpLessOrEqual    = "<="
	RelOpGreaterOrEqual = ">="
)

//=============================================================================
//===
//=== Relational
//===
//=============================================================================

type RelationalExpression struct {
	operand string
	left    Expression
	right   Expression
	info    *parser.Info
}

//=============================================================================

func NewRelationalExpression(operand string, left, right Expression, info *parser.Info) *RelationalExpression {
	return &RelationalExpression{
		operand: operand,
		left:    left,
		right:   right,
		info:    info,
	}
}

//=============================================================================

func (e *RelationalExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t1, err1 := e.left.ResolveType(scope, embedder, depth)
	t2, err2 := e.right.ResolveType(scope, embedder, depth)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	ok := false

	switch e.operand {
	case RelOpEqual:
		ok = canEquateTo(t1, t2)
	case RelOpNotEqual:
		ok = canEquateTo(t1, t2)
	case RelOpLessThan:
		ok = canCompareTo(t1, t2)
	case RelOpLessOrEqual:
		ok = canCompareTo(t1, t2)
	case RelOpGreaterThan:
		ok = canCompareTo(t1, t2)
	case RelOpGreaterOrEqual:
		ok = canCompareTo(t1, t2)
	}

	if !ok {
		return nil, errors.New("operand '" + e.operand + "'s not usable with '" + t1.String() + "' and '" + t2.String() + "'")
	}

	return types.NewBoolType(), nil
}

//=============================================================================

func (e *RelationalExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func canEquateTo(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	if t1.Code() == t2.Code() {
		return true
	}

	if t1.Code() == types.CodeInt && t2.Code() == types.CodeReal {
		return true
	}

	if t1.Code() == types.CodeReal && t2.Code() == types.CodeInt {
		return true
	}

	return false
}

//=============================================================================

func canCompareTo(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	if t1.Code() == types.CodeDate && t2.Code() == types.CodeDate {
		return true
	}

	if t1.Code() == types.CodeTime && t2.Code() == types.CodeTime {
		return true
	}

	if t1.Code() == types.CodeString && t2.Code() == types.CodeString {
		return true
	}

	if (t1.Code() == types.CodeInt || t1.Code() == types.CodeReal) && (t2.Code() == types.CodeInt || t2.Code() == types.CodeReal) {
		return true
	}

	return false
}

//=============================================================================
