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
	"github.com/algotiqa/tiq-engine/core/types"
)

//=============================================================================
//===
//=== Script
//===
//=============================================================================

type Script struct {
	Filename    string
	PackageName string
	Constants   []*Constant
	Variables   []*Variable
	Functions   []*types.Function
	Enums       []*types.EnumType
	Classes     []*types.ClassType
}

//=============================================================================

func NewScript(filename string) *Script {
	return &Script{
		Filename: filename,
	}
}

//=============================================================================

func (s *Script) AddConstant(c *Constant) {
	s.Constants = append(s.Constants, c)
}

//=============================================================================

func (s *Script) AddVariable(v *Variable) {
	s.Variables = append(s.Variables, v)
}

//=============================================================================

func (s *Script) AddFunction(f *types.Function) {
	s.Functions = append(s.Functions, f)
}

//=============================================================================

func (s *Script) AddEnum(e *types.EnumType) {
	s.Enums = append(s.Enums, e)
}

//=============================================================================

func (s *Script) AddClass(c *types.ClassType) {
	s.Classes = append(s.Classes, c)
}

//=============================================================================
