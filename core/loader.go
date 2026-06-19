//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package core

import (
	"io"

	"github.com/algotiqa/tiq-engine/ast"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/tiq-engine/runtime"
	"github.com/algotiqa/tiq-engine/tool"
	"github.com/antlr4-go/antlr/v4"
)

//=============================================================================
//===
//=== Environment creation
//===
//=============================================================================

func CreateEnvironment(path string) (*runtime.Environment, *parser.ParseErrors) {
	pr := parser.NewParseErrors()

	files, err := tool.FindFiles(path, ".tsl")
	if err != nil {
		pe := parser.NewParseError(path, -1, -1, err.Error())
		pr.AddError(pe)
		return nil, pr
	}

	e := runtime.NewEnvironment()

	for _, file := range files {
		script, errs := ParseFile(file)
		if !errs.IsEmpty() {
			return nil, errs
		}

		errs = e.AddScript(script)
		if !errs.IsEmpty() {
			return nil, errs
		}
	}

	return e, pr
}

//=============================================================================
//===
//=== Parsing functions
//===
//=============================================================================

func ParseFile(filename string) (*ast.Script, *parser.ParseErrors) {
	stream, err := antlr.NewFileStream(filename)
	if err != nil {
		pr := parser.NewParseErrors()
		pe := parser.NewParseError(filename, -1, -1, err.Error())
		pr.AddError(pe)
		return nil, pr
	}

	return parse(stream)
}

//=============================================================================

func ParseString(input string) (*ast.Script, *parser.ParseErrors) {
	return parse(antlr.NewInputStream(input))
}

//=============================================================================

func ParseReader(r io.Reader) (*ast.Script, *parser.ParseErrors) {
	return parse(antlr.NewIoStream(r))
}

//=============================================================================

func parse(input antlr.CharStream) (*ast.Script, *parser.ParseErrors) {
	lexer := parser.NewTiqLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewTiqParser(stream)

	res := parser.NewParseErrors()
	lis := parser.NewParseErrorListener(input.GetSourceName(), res)

	//--- Add error collection listener
	p.RemoveErrorListeners()
	p.AddErrorListener(lis)

	//--- Ask the parser to report all ambiguities
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	tree := p.Script()
	if !res.IsEmpty() {
		return nil, res
	}

	script := ast.Build(tree)

	return script, res
}

//=============================================================================
