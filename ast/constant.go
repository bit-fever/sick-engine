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
//=== Constant
//===
//=============================================================================

type Constant struct {
	name  string
	Value expression.Expression
	Type  types.Type
	Info  *parser.Info
}

//=============================================================================

func NewConstant(name string, e expression.Expression, info *parser.Info) *Constant {
	return &Constant{
		name:  name,
		Value: e,
		Info:  info,
	}
}

//=============================================================================

func (c *Constant) SetupType(s interfaces.Scope, depth int) error {
	t, err := c.Value.ResolveType(s, nil, depth)
	if err != nil {
		return err
	}

	c.Type = t
	return nil
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (c *Constant) Id() string {
	return c.name
}

//=============================================================================

func (c *Constant) Kind() interfaces.Kind {
	return interfaces.KindConst
}

//=============================================================================

func (c *Constant) Specie() interfaces.Specie {
	return interfaces.SpecieObject
}

//=============================================================================

func (c *Constant) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (c *Constant) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
