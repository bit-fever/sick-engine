//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package interfaces

import "github.com/algotiqa/tiq-engine/parser"

//=============================================================================

type Kind int

const (
	KindConst Kind = iota
	KindVar
	KindFunction
	KindClass
	KindEnum
	KindPackage
	KindEnumItem
	KindParameter
	KindProperty
)

//=============================================================================

type Specie int

const (
	SpecieObject Specie = iota
	SpecieType
	SpecieOther
)

//=============================================================================

type Symbol interface {
	Id() string
	Kind() Kind
	Specie() Specie

	InitScope(parent Scope) *parser.ParseError
	Scope() Scope
}

//=============================================================================
