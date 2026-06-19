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
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== FQIdentifier
//===
//=============================================================================

type FQIdentifier struct {
	Pack string
	Name string
}

//=============================================================================

func NewFQIdentifier(tree parser.IFqIdentifierContext) *FQIdentifier {
	if tree.DOT() == nil {
		return &FQIdentifier{
			Pack: "",
			Name: tree.IDENTIFIER(0).GetText(),
		}
	}

	return &FQIdentifier{
		Pack: tree.IDENTIFIER(0).GetText(),
		Name: tree.IDENTIFIER(1).GetText(),
	}
}

//=============================================================================

func (i *FQIdentifier) String() string {
	if i.Pack == "" {
		return i.Name
	}

	return i.Pack + "." + i.Name
}

//=============================================================================
