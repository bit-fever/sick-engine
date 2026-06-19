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
	"strings"

	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Function
//===
//=============================================================================

type Function struct {
	name    string
	Class   *data.FQIdentifier
	Params  []*Param
	Returns []Type
	Block   *data.Block
	Info    *parser.Info
	scope   interfaces.Scope
}

//=============================================================================

func NewFunction(name string, info *parser.Info) *Function {
	return &Function{
		name: name,
		Info: info,
	}
}

//=============================================================================

func (f *Function) AddParam(p *Param) {
	f.Params = append(f.Params, p)
}

//=============================================================================

func (f *Function) AddReturnType(t Type) {
	f.Returns = append(f.Returns, t)
}

//=============================================================================

func (f *Function) SetEmbedder(embedder interfaces.Symbol) {
	f.Block.SetEmbedder(embedder)
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (f *Function) Id() string {
	sb := strings.Builder{}
	sb.WriteString(f.name)
	sb.WriteString("|")

	for _, p := range f.Params {
		sb.WriteString("|")
		sb.WriteString(p.Type.String())
	}

	return sb.String()
}

//=============================================================================

func (f *Function) Kind() interfaces.Kind {
	return interfaces.KindFunction
}

//=============================================================================

func (f *Function) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (f *Function) InitScope(parent interfaces.Scope) *parser.ParseError {
	f.scope = parent.Push()

	for _, param := range f.Params {
		if !f.scope.Define(param) {
			return parser.NewParseErrorFromInfo(param.Info, "parameter duplicated in function: "+param.Id())
		}
	}

	return f.Block.InitScope(f.scope)
}

//=============================================================================

func (f *Function) Scope() interfaces.Scope {
	return f.scope
}

//=============================================================================
//===
//=== Param
//===
//=============================================================================

type Param struct {
	Name string
	Type Type
	Info *parser.Info
}

//=============================================================================

func NewParam(name string, t Type, info *parser.Info) *Param {
	return &Param{
		Name: name,
		Type: t,
		Info: info,
	}
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Param) Id() string {
	return p.Name
}

//=============================================================================

func (p *Param) Kind() interfaces.Kind {
	return interfaces.KindParameter
}

//=============================================================================

func (p *Param) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================

func (p *Param) InitScope(parent interfaces.Scope) *parser.ParseError {
	return nil
}

//=============================================================================

func (p *Param) Scope() interfaces.Scope {
	return nil
}

//=============================================================================
