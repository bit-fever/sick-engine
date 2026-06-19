//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

import (
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Enum
//===
//=============================================================================

type EnumType struct {
	name  string
	IsInt bool
	Items []*EnumItem
	Info  *parser.Info
	scope interfaces.Scope
}

//=============================================================================

func NewEnumType(name string, isInt bool, info *parser.Info) *EnumType {
	return &EnumType{
		name:  name,
		IsInt: isInt,
		Info:  info,
	}
}

//=============================================================================

func (e *EnumType) AddItem(i *EnumItem) {
	e.Items = append(e.Items, i)
}

//=============================================================================

func (e *EnumType) Size() int {
	return len(e.Items)
}

//=============================================================================

func (e *EnumType) AssignCodes() {
	for i, item := range e.Items {
		item.Code = i + 1
	}
}

//=============================================================================

func (e *EnumType) Code() int8 {
	return CodeEnum
}

//=============================================================================

func (e *EnumType) String() string {
	return "enum=" + e.name
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (e *EnumType) Id() string {
	return e.name
}

//=============================================================================

func (e *EnumType) Kind() interfaces.Kind {
	return interfaces.KindEnum
}

//=============================================================================

func (e *EnumType) Specie() interfaces.Specie {
	return interfaces.SpecieType
}

//=============================================================================

func (e *EnumType) InitScope(parent interfaces.Scope) *parser.ParseError {
	e.scope = data.NewSymbolTable()

	for _, item := range e.Items {
		if !e.scope.Define(item) {
			return parser.NewParseErrorFromInfo(e.Info, "item duplicated in enum: "+item.Id())
		}
	}
	return nil
}

//=============================================================================

func (e *EnumType) Scope() interfaces.Scope {
	return e.scope
}

//=============================================================================
//===
//=== EnumItem
//===
//=============================================================================

type EnumItem struct {
	name  string
	Code  int
	Value string
}

//=============================================================================

func NewEnumItem(name string, code int, value string) *EnumItem {
	return &EnumItem{name, code, value}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (e *EnumItem) Id() string {
	return e.name
}

//=============================================================================

func (e *EnumItem) Kind() interfaces.Kind {
	return interfaces.KindEnumItem
}

//=============================================================================

func (e *EnumItem) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (e *EnumItem) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (e *EnumItem) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
