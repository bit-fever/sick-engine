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
	"strconv"

	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/values"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/types"
	"github.com/antlr4-go/antlr/v4"
)

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

func ConvertExpression(tree parser.IExpressionContext) Expression {
	if tree == nil {
		return nil
	}

	//--- First, try with an arithmetic expression

	ae := convertArithmeticExpression(tree)
	if ae != nil {
		return ae
	}

	//--- Then try with a relational expression

	re := convertRelationalExpression(tree)
	if re != nil {
		return re
	}

	//--- Now, try with AND/OR

	be := convertBooleanExpression(tree)
	if be != nil {
		return be
	}

	//--- Lastly, other possibilities

	unaryExpr := tree.UnaryExpression()
	listExpr := tree.ListExpression()
	mapExpr := tree.MapExpression()

	if unaryExpr != nil {
		return convertUnaryExpression(unaryExpr)
	} else if listExpr != nil {
		return convertListExpression(listExpr)
	} else if mapExpr != nil {
		return convertMapExpression(mapExpr)
	} else {
		panic("Unknown expression type : " + tree.GetText())
	}
}

//=============================================================================

func convertArithmeticExpression(tree parser.IExpressionContext) Expression {
	if tree.STAR() != nil || tree.SLASH() != nil || tree.PLUS() != nil || tree.MINUS() != nil {
		expr1 := ConvertExpression(tree.Expression(0))
		expr2 := ConvertExpression(tree.Expression(1))
		info := parser.NewInfo(tree)

		if tree.STAR() != nil {
			return NewArithmeticExpression(AritOpMult, expr1, expr2, info)
		}

		if tree.SLASH() != nil {
			return NewArithmeticExpression(AritOpDiv, expr1, expr2, info)
		}

		if tree.PLUS() != nil {
			return NewArithmeticExpression(AritOpAdd, expr1, expr2, info)
		}

		if tree.MINUS() != nil {
			return NewArithmeticExpression(AritOpSub, expr1, expr2, info)
		}
	}

	return nil
}

//=============================================================================

func convertRelationalExpression(tree parser.IExpressionContext) Expression {
	if tree.EQUAL() != nil || tree.NOT_EQUAL() != nil ||
		tree.LESS_THAN() != nil || tree.GREATER_THAN() != nil ||
		tree.LESS_OR_EQUAL() != nil || tree.GREATER_OR_EQUAL() != nil {

		expr1 := ConvertExpression(tree.Expression(0))
		expr2 := ConvertExpression(tree.Expression(1))
		info := parser.NewInfo(tree)

		if tree.EQUAL() != nil {
			return NewRelationalExpression(RelOpEqual, expr1, expr2, info)
		}

		if tree.LESS_OR_EQUAL() != nil {
			return NewRelationalExpression(RelOpLessOrEqual, expr1, expr2, info)
		}

		if tree.GREATER_OR_EQUAL() != nil {
			return NewRelationalExpression(RelOpGreaterOrEqual, expr1, expr2, info)
		}

		if tree.LESS_THAN() != nil {
			return NewRelationalExpression(RelOpLessThan, expr1, expr2, info)
		}

		if tree.GREATER_THAN() != nil {
			return NewRelationalExpression(RelOpGreaterThan, expr1, expr2, info)
		}

		if tree.NOT_EQUAL() != nil {
			return NewRelationalExpression(RelOpNotEqual, expr1, expr2, info)
		}
	}

	return nil
}

//=============================================================================

func convertBooleanExpression(tree parser.IExpressionContext) Expression {
	if tree.AllAND() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return NewAndExpression(expr, parser.NewInfo(tree))
	}

	if tree.AllOR() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return NewOrExpression(expr, parser.NewInfo(tree))
	}

	if tree.NOT() != nil {
		expr := ConvertExpression(tree.Expression(0))
		return NewNotExpression(expr, parser.NewInfo(tree))
	}

	return nil
}

//=============================================================================

func buildExpressionList(list []parser.IExpressionContext) []Expression {
	var res []Expression

	for _, expr := range list {
		res = append(res, ConvertExpression(expr))
	}

	return res
}

//=============================================================================

func convertUnaryExpression(tree parser.IUnaryExpressionContext) Expression {
	exprParen := tree.ExpressionInParenthesis()
	constValue := tree.ConstantValueExpression()
	chainExpr := tree.ChainedExpression()
	barExpr := tree.BarExpression()

	if exprParen != nil {
		return ConvertExpression(exprParen.Expression())
	} else if constValue != nil {
		return convertConstantValueExpression(constValue)
	} else if chainExpr != nil {
		return ConvertChainedExpression(chainExpr)
	} else if barExpr != nil {
		return convertBarAccessExpression(barExpr)
	}

	if tree.PLUS() != nil {
		return convertUnaryExpression(tree.UnaryExpression())
	}

	if tree.MINUS() != nil {
		value := values.NewIntValue(int64(-1))
		expr1 := NewConstantValueExpression(value, parser.NewInfo(tree))
		expr2 := convertUnaryExpression(tree.UnaryExpression())
		return NewArithmeticExpression(AritOpMult, expr1, expr2, parser.NewInfo(tree))
	}

	panic("Unknown unary expression type : " + tree.GetText())
}

//=============================================================================

func convertBarAccessExpression(tree parser.IBarExpressionContext) Expression {
	var acc Expression

	accessor := tree.AccessorExpression()
	if accessor != nil {
		acc = ConvertExpression(accessor.Expression())
	}

	bar := 0

	if tree.OPEN() != nil {
		bar = BarOpen
	} else if tree.HIGH() != nil {
		bar = BarHigh
	} else if tree.LOW() != nil {
		bar = BarLow
	} else if tree.CLOSE() != nil {
		bar = BarClose
	} else {
		panic("Unknown bar expression type : " + tree.GetText())
	}

	return NewBarAccessExpression(bar, acc, parser.NewInfo(tree))
}

//=============================================================================

func ConvertChainedExpression(tree parser.IChainedExpressionContext) *ChainedExpression {
	ce := NewChainedExpression(tree.THIS() != nil, tree.NEW() != nil, parser.NewInfo(tree))

	if tree.NEW() != nil {
		ce.FQClass = data.NewFQIdentifier(tree.FqIdentifier())
		ce.InstParams = buildExpressionList(tree.ParamsExpression().AllExpression())
	}

	for _, item := range tree.AllChainItem() {
		ci := convertChainItem(item)
		ce.AddItem(ci)
	}

	return ce
}

//=============================================================================

func convertChainItem(tree parser.IChainItemContext) *ChainItem {
	name := tree.IDENTIFIER().GetText()
	params := tree.ParamsExpression()
	acc := tree.AccessorExpression()

	var accessor Expression
	var paramList []Expression

	if acc != nil {
		accessor = ConvertExpression(acc.Expression())
	}

	if params != nil {
		paramList = buildExpressionList(params.AllExpression())
	}

	return NewChainItem(name, accessor, paramList)
}

//=============================================================================

func convertListExpression(tree parser.IListExpressionContext) Expression {
	le := NewListExpression(parser.NewInfo(tree))

	for _, e := range tree.AllExpression() {
		le.AddExpression(ConvertExpression(e))
	}

	return le
}

//=============================================================================

func convertMapExpression(tree parser.IMapExpressionContext) Expression {
	mex := NewMapExpression(parser.NewInfo(tree))

	for _, me := range tree.AllKeyValueCouple() {
		k := convertKeyValue(me.KeyValue())
		e := ConvertExpression(me.Expression())
		mex.Set(k, e)
	}

	return mex
}

//=============================================================================

func convertKeyValue(tree parser.IKeyValueContext) values.Value {
	//--- Integers

	intVal := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	panic("Unknown constant expression type : " + tree.GetText())
}

//=============================================================================
//===
//=== Constant values
//===
//=============================================================================

func convertConstantValueExpression(tree parser.IConstantValueExpressionContext) Expression {
	value := convertConstantValue(tree)
	return NewConstantValueExpression(value, parser.NewInfo(tree))
}

//=============================================================================

func convertConstantValue(tree parser.IConstantValueExpressionContext) values.Value {
	//--- Integers

	intVal := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Reals

	realVal := tree.REAL_VALUE()
	if realVal != nil {
		return convertConstantReal(tree.GetParser(), realVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	//--- Booleans

	boolVal := tree.BoolValue()
	if boolVal != nil {
		v := boolVal.GetText() == "true"

		return values.NewBoolValue(v)
	}

	//--- Errors

	errVal := tree.ErrorValue()
	if errVal != nil {
		return convertConstantError(errVal.STRING_VALUE())
	}

	//--- Unknown

	panic("Unknown constant expression type : " + tree.GetText())
	return nil
}

//=============================================================================

func convertConstantInt(p antlr.Parser, val antlr.TerminalNode) values.Value {
	i, err := strconv.Atoi(val.GetText())
	if err != nil {
		parser.RaiseError(p, "invalid integer value : "+val.GetText())
	}

	return values.NewIntValue(int64(i))
}

//=============================================================================

func convertConstantReal(p antlr.Parser, val antlr.TerminalNode) values.Value {
	f, err := strconv.ParseFloat(val.GetText(), 64)
	if err != nil {
		parser.RaiseError(p, "Invalid real value : "+val.GetText())
	}

	return values.NewRealValue(f)
}

//=============================================================================

func convertConstantTime(p antlr.Parser, val parser.ITimeValueContext) values.Value {
	hh, _ := strconv.Atoi(val.INT_VALUE(0).GetText())
	mm, _ := strconv.Atoi(val.INT_VALUE(1).GetText())

	v := types.NewTime(hh, mm)

	if !v.IsValid() {
		parser.RaiseError(p, "Invalid time value : "+val.GetText())
	}

	return values.NewTimeValue(v)
}

//=============================================================================

func convertConstantDate(p antlr.Parser, val parser.IDateValueContext) values.Value {
	y, _ := strconv.Atoi(val.INT_VALUE(0).GetText())
	m, _ := strconv.Atoi(val.INT_VALUE(1).GetText())
	d, _ := strconv.Atoi(val.INT_VALUE(2).GetText())

	v := types.NewDate(y, m, d)

	if !v.IsValid() {
		parser.RaiseError(p, "Invalid date value : "+val.GetText())
	}

	return values.NewDateValue(v)
}

//=============================================================================

func convertConstantString(val antlr.TerminalNode) values.Value {
	text := val.GetText()
	text = text[1 : len(text)-1]

	return values.NewStringValue(text)
}

//=============================================================================

func convertConstantError(val antlr.TerminalNode) values.Value {
	text := val.GetText()
	text = text[1 : len(text)-1]

	return values.NewErrorValue(text)
}

//=============================================================================
