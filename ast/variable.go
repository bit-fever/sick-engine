//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package ast

import (
	"github.com/algotiqa/tiq-engine/ast/expression"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Variable
//===
//=============================================================================

type Variable struct {
	name  string
	Value expression.Expression
	Type  types.Type
	Info  *parser.Info
}

//=============================================================================

func NewVariable(name string, e expression.Expression, info *parser.Info) *Variable {
	return &Variable{
		name:  name,
		Value: e,
		Info:  info,
	}
}

//=============================================================================

func (v *Variable) SetupType(s interfaces.Scope, depth int) error {
	t, err := v.Value.ResolveType(s, nil, depth)
	if err != nil {
		return err
	}

	v.Type = t
	return nil
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (v *Variable) Id() string {
	return v.name
}

//=============================================================================

func (v *Variable) Kind() interfaces.Kind {
	return interfaces.KindConst
}

//=============================================================================

func (v *Variable) Specie() interfaces.Specie {
	return interfaces.SpecieObject
}

//=============================================================================

func (v *Variable) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (v *Variable) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
