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
//=== Class
//===
//=============================================================================

type ClassType struct {
	name       string
	Properties []*Property
	Functions  []*Function
	Info       *parser.Info
	classScope interfaces.Scope
}

//=============================================================================

func NewClassType(name string, info *parser.Info) *ClassType {
	return &ClassType{
		name: name,
		Info: info,
	}
}

//=============================================================================

func (t *ClassType) AddProperty(p *Property) {
	t.Properties = append(t.Properties, p)
}

//=============================================================================

func (t *ClassType) AddFunction(f *Function) {
	t.Functions = append(t.Functions, f)
}

//=============================================================================

func (t *ClassType) Code() int8 {
	return CodeClass
}

//=============================================================================

func (t *ClassType) String() string {
	return "class=(" + t.name + ")"
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (t *ClassType) Id() string {
	return t.name
}

//=============================================================================

func (t *ClassType) Kind() interfaces.Kind {
	return interfaces.KindClass
}

//=============================================================================

func (t *ClassType) Specie() interfaces.Specie {
	return interfaces.SpecieType
}

//=============================================================================

func (t *ClassType) InitScope(parent interfaces.Scope) *parser.ParseError {
	t.classScope = data.NewSymbolTable()

	for _, prop := range t.Properties {
		if !t.classScope.Define(prop) {
			return parser.NewParseErrorFromInfo(prop.Info, "property duplicated in class: "+prop.Id())
		}
	}

	for _, f := range t.Functions {
		if !t.classScope.Define(f) {
			return parser.NewParseErrorFromInfo(f.Info, "function duplicated in class: "+f.Id())
		}

		err := f.InitScope(parent)
		if err != nil {
			return err
		}

		f.SetEmbedder(t)
	}

	return nil
}

//=============================================================================

func (t *ClassType) Scope() interfaces.Scope {
	return t.classScope
}

//=============================================================================
//===
//=== Property
//===
//=============================================================================

type Property struct {
	name string
	Type Type
	Info *parser.Info
}

//=============================================================================

func NewProperty(name string, type_ Type, info *parser.Info) *Property {
	return &Property{
		name: name,
		Type: type_,
		Info: info,
	}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Property) Id() string {
	return p.name
}

//=============================================================================

func (p *Property) Kind() interfaces.Kind {
	return interfaces.KindProperty
}

//=============================================================================

func (p *Property) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (p *Property) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (p *Property) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
