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
	"fmt"
	"strings"
)

//=============================================================================
//===
//=== Parse error
//===
//=============================================================================

type ParseErrors struct {
	Errors []*ParseError
}

//=============================================================================

func NewParseErrors() *ParseErrors {
	return &ParseErrors{
	}
}

//=============================================================================

func (pe *ParseErrors) AddError(p *ParseError) {
	pe.Errors = append(pe.Errors, p)
}

//=============================================================================

func (pe *ParseErrors) IsEmpty() bool {
	return len(pe.Errors) == 0
}

//=============================================================================

func (pe *ParseErrors) String() string {
	var sb strings.Builder

	for _, err := range pe.Errors {
		sb.WriteString(err.String() + "\n")
	}

	return sb.String()
}

//=============================================================================
//===
//=== Parse error
//===
//=============================================================================

type ParseError struct {
	File   string
	Line   int
	Column int
	Error  string
}

//=============================================================================

func NewParseError(file string, line int, column int, err string) *ParseError {
	return &ParseError{file, line, column, err}
}

//=============================================================================

func NewParseErrorFromInfo(info *Info, err string) *ParseError {
	return &ParseError{info.Filename, info.Line, info.Column, err}
}

//=============================================================================

func (pe *ParseError) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("File: %s, Line: %d, Col: %d, Error: %s", pe.File, pe.Line, pe.Column, pe.Error))
	return sb.String()
}

//=============================================================================
