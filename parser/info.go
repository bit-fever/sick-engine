//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

//=============================================================================

type Info struct {
	Filename string
	Line     int
	Column   int
	Fragment string
}

//=============================================================================

func NewInfo(tree antlr.ParserRuleContext) *Info {
	token := tree.GetStart()

	return &Info{
		Filename: token.GetInputStream().GetSourceName(),
		Line    : token.GetLine(),
		Column  : token.GetColumn(),
		Fragment: tree.GetText(),
	}
}

//=============================================================================
