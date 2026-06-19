//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package data

import (
	"iter"
	"maps"

	"github.com/algotiqa/tiq-engine/core/interfaces"
)

//=============================================================================
//===
//=== Scope implementation
//===
//=============================================================================

type SymbolTable struct {
	previous *SymbolTable
	symbols  map[string]interfaces.Symbol
}

//=============================================================================

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]interfaces.Symbol),
	}
}

//=============================================================================

func (s *SymbolTable) Resolve(name string) interfaces.Symbol {
	currScope := s

	for currScope != nil {
		sym, ok := currScope.symbols[name]
		if ok {
			return sym
		}

		//if sym == nil {
		//	//--- Functions have Ids (not names) that must be processed differently
		//	for k,v := range s.symbols {
		//		if strings.HasPrefix(k,name+"|") {
		//			return v
		//		}
		//	}
		//}
		currScope = currScope.previous
	}

	return nil
}

//=============================================================================

func (s *SymbolTable) Define(symbol interfaces.Symbol) bool {
	_, ok := s.symbols[symbol.Id()]
	if ok {
		return false
	}

	s.symbols[symbol.Id()] = symbol
	return true
}

//=============================================================================

func (s *SymbolTable) Push() interfaces.Scope {
	ns := NewSymbolTable()
	ns.previous = s
	return ns
}

//=============================================================================

func (s *SymbolTable) Pop() interfaces.Scope {
	return s.previous
}

//=============================================================================

func (s *SymbolTable) AllSymbols() iter.Seq[interfaces.Symbol] {
	return maps.Values(s.symbols)
}

//=============================================================================
