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
//===
//=== Error listener
//===
//=============================================================================

type ParseErrorListener struct {
	filename string
	errors   *ParseErrors
}

//=============================================================================

func NewParseErrorListener(filename string, errors *ParseErrors) *ParseErrorListener {
	return &ParseErrorListener{
		filename: filename,
		errors  : errors,
	}
}

//=============================================================================

func (l *ParseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	pe := NewParseError(l.filename, line, column, msg)
	l.errors.AddError(pe)
}

//=============================================================================

func (l *ParseErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

//=============================================================================

func (l *ParseErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

//=============================================================================

func (l *ParseErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

//=============================================================================
//===
//=== Raise error general functions
//===
//=============================================================================

func RaiseError(p antlr.Parser, message string) {
	p.NotifyErrorListeners(message, nil, nil)
}

//=============================================================================
