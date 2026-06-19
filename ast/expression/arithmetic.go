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
	AritOpAdd  = "+"
	AritOpSub  = "-"
	AritOpMult = "*"
	AritOpDiv  = "/"
)

//=============================================================================
//===
//=== Arithmetic
//===
//=============================================================================

type ArithmeticExpression struct {
	operand string
	left    Expression
	right   Expression
	info    *parser.Info
}

//=============================================================================

func NewArithmeticExpression(operand string, left, right Expression, info *parser.Info) *ArithmeticExpression {
	return &ArithmeticExpression{
		operand: operand,
		left:    left,
		right:   right,
		info:    info,
	}
}

//=============================================================================

func (e *ArithmeticExpression) ResolveType(scope interfaces.Scope, embedder interfaces.Symbol, depth int) (types.Type, error) {
	t1, err1 := e.left.ResolveType(scope, embedder, depth)
	t2, err2 := e.right.ResolveType(scope, embedder, depth)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	if supportsStringify(t1, t2) {
		return types.NewStringType(), nil
	}

	t := combineTypes(e.operand, t1, t2)
	if t == nil {
		return nil, errors.New("operand types mismatch: '" + t1.String() + "' and '" + t2.String() + "'")
	}

	return t, nil
}

//=============================================================================

func (e *ArithmeticExpression) Info() *parser.Info {
	return e.info
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func supportsStringify(t1, t2 types.Type) bool {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return false
	}

	return t1.Code() == types.CodeString || t2.Code() == types.CodeString
}

//=============================================================================

func combineTypes(operand string, t1, t2 types.Type) types.Type {
	if t1.Code() == types.CodeTimeseries || t2.Code() == types.CodeTimeseries {
		return nil
	}

	if t1.Code() == types.CodeInt && t2.Code() == types.CodeInt {
		return types.NewIntType()
	}

	if (t1.Code() == types.CodeInt || t1.Code() == types.CodeReal) && (t2.Code() == types.CodeInt || t2.Code() == types.CodeReal) {
		return types.NewRealType()
	}

	if operand == AritOpAdd || operand == AritOpSub {
		if (t1.Code() == types.CodeTime || t1.Code() == types.CodeInt) && (t2.Code() == types.CodeTime || t2.Code() == types.CodeInt) {
			return types.NewTimeType()
		}

		if (t1.Code() == types.CodeDate && t2.Code() == types.CodeInt) || (t1.Code() == types.CodeInt && t2.Code() == types.CodeDate) {
			return types.NewDateType()
		}
	}

	return nil
}

//=============================================================================
