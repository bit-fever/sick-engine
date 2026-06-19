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
	"github.com/algotiqa/tiq-engine/core/interfaces"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Statement block
//===
//=============================================================================

type Block struct {
	Statements []interfaces.Statement
	scope      interfaces.Scope
	embedder   interfaces.Symbol
}

//=============================================================================

func (b *Block) Add(s interfaces.Statement) {
	b.Statements = append(b.Statements, s)
}

//=============================================================================

func (b *Block) InitScope(parent interfaces.Scope) *parser.ParseError {
	b.scope = parent.Push()
	return nil
}

//=============================================================================

func (b *Block) SetEmbedder(embedder interfaces.Symbol) {
	b.embedder = embedder
}

//=============================================================================
