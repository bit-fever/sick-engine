//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package runtime

import (
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Package
//===
//=============================================================================

type Package struct {
	name      string
	scope     interfaces.Scope
	classFunc []*types.Function
}

//=============================================================================

func NewPackage(name string, parent interfaces.Scope) *Package {
	return &Package{
		name:  name,
		scope: parent.Push(),
	}
}

//=============================================================================

func (p *Package) AddFunction(f *types.Function) {
	f.Class.Pack = p.name
	p.classFunc = append(p.classFunc, f)
}

//=============================================================================

func (p *Package) AssignMethodsToClasses() *parser.ParseError {
	for _, f := range p.classFunc {
		s := p.scope.Resolve(f.Class.Name)
		if s != nil {
			if s.Kind() != interfaces.KindClass {
				return parser.NewParseErrorFromInfo(f.Info, "function doesn't reference a class: "+f.Class.Name)
			}

			c := s.(*types.ClassType)
			c.AddFunction(f)
		} else {
			return parser.NewParseErrorFromInfo(f.Info, "can't resolve class for function: "+f.Class.String())
		}
	}

	p.classFunc = nil
	return nil
}

//=============================================================================
//=== Symbol interface
//=============================================================================

func (p *Package) Id() string {
	return p.name
}

//=============================================================================

func (p *Package) Kind() interfaces.Kind {
	return interfaces.KindPackage
}

//=============================================================================

func (p *Package) Specie() interfaces.Specie {
	return interfaces.SpecieOther
}

//=============================================================================
//--- A package already has a scope populated with data
//--- We just need to initialize the scopes of the children

func (p *Package) InitScope(parent interfaces.Scope) *parser.ParseError {
	for s := range p.scope.AllSymbols() {
		err := s.InitScope(p.scope)
		if err != nil {
			return err
		}
	}

	return nil
}

//=============================================================================

func (p *Package) Scope() interfaces.Scope {
	return p.scope
}

//=============================================================================
