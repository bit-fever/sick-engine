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
	"strconv"

	"github.com/algotiqa/tiq-engine/ast/expression"
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/tiq-engine/tool"
)

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

func ConvertBlock(tree parser.IBlockContext) *data.Block {
	block := &data.Block{}

	for _, stmt := range tree.AllStatement() {
		varsDecl := stmt.VarsDeclaration()
		varsAss := stmt.VarsAssignmentOrFunctionCall()
		ifStmt := stmt.IfStatement()
		retStmt := stmt.ReturnStatement()

		if varsDecl != nil {
			for _, va := range varsDecl.AllVarDeclaration() {
				block.Add(convertVarDeclaration(va))
			}
		} else if varsAss != nil {
			for _, s := range convertVarAssignmentOrFunctionCall(varsAss) {
				block.Add(s)
			}
		} else if ifStmt != nil {
			block.Add(convertIfStatement(ifStmt))
		} else if retStmt != nil {
			block.Add(convertReturnStatement(retStmt))
		} else {
			panic("Unknown statement type : " + stmt.GetText())
		}
	}

	return block
}

//=============================================================================

func convertVarDeclaration(tree parser.IVarDeclarationContext) *VarDeclaration {
	name := tree.IDENTIFIER().GetText()
	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "Variables must be in lowercase: "+name)
	}

	var type_ types.Type
	var expr expression.Expression

	if tree.Type_() != nil {
		type_ = types.ConvertType(tree.Type_())
	}

	if tree.Expression() != nil {
		expr = expression.ConvertExpression(tree.Expression())
	}

	if type_ == nil && expr == nil {
		parser.RaiseError(tree.GetParser(), "variable declaration requires either type or value: "+name)
	}

	return NewVarDeclaration(name, type_, expr)
}

//=============================================================================

func convertVarAssignmentOrFunctionCall(tree parser.IVarsAssignmentOrFunctionCallContext) []interfaces.Statement {
	var chains []*expression.ChainedExpression
	var expressions []expression.Expression

	//--- Collect identifiers

	for _, chainExpr := range tree.AllChainedExpression() {
		ce := expression.ConvertChainedExpression(chainExpr)
		chains = append(chains, ce)
	}

	//--- Collect expressions

	for _, e := range tree.AllExpression() {
		expr := expression.ConvertExpression(e)
		expressions = append(expressions, expr)
	}

	//--- Case 1: No EQUAL sign
	//--- Just a sequence of chain expressions. These must be function calls

	var vas []interfaces.Statement

	if tree.EQUAL() == nil {
		for _, chain := range chains {
			va := NewFunctionCallStatement(chain)
			vas = append(vas, va)
		}

		return vas
	}

	//--- Case 2: we can split into several simple assignments

	if len(chains) == len(expressions) {
		for i, chain := range chains {
			va := NewVarAssignment(chain, expressions[i])
			vas = append(vas, va)
		}

		return vas
	}

	//--- Case 3: several identifiers take the result of a function call

	if len(expressions) != 1 {
		parser.RaiseError(tree.GetParser(), "identifiers don't match expressions: chains="+
			strconv.Itoa(len(chains))+", expr="+
			strconv.Itoa(len(expressions)))
	}

	expr := expressions[0]
	switch expr.(type) {
	case *expression.ChainedExpression:
		fc := expr.(*expression.ChainedExpression)
		vas = append(vas, NewFunctionCallAssignment(chains, fc))
		return vas
	}

	//--- Case 4: same expression cloned to several identifiers

	for _, chain := range chains {
		va := NewVarAssignment(chain, expressions[0])
		vas = append(vas, va)
	}

	return vas
}

//=============================================================================

func convertIfStatement(tree parser.IIfStatementContext) *IfStatement {
	is := NewIfStatement()
	cond := expression.ConvertExpression(tree.Expression())
	block := ConvertBlock(tree.Block())
	is.AddConditionalBlock(NewConditionalBlock(cond, block))

	for _, eib := range tree.AllElseIfBlock() {
		cond = expression.ConvertExpression(eib.Expression())
		block = ConvertBlock(eib.Block())

		is.AddConditionalBlock(NewConditionalBlock(cond, block))
	}

	if tree.ElseBlock() != nil {
		cond = nil
		block = ConvertBlock(tree.ElseBlock().Block())

		is.AddConditionalBlock(NewConditionalBlock(cond, block))
	}

	return is
}

//=============================================================================

func convertReturnStatement(tree parser.IReturnStatementContext) *ReturnStatement {
	rs := NewReturnStatement()

	for _, e := range tree.AllExpression() {
		expr := expression.ConvertExpression(e)
		rs.AddExpression(expr)
	}

	return rs
}

//=============================================================================
