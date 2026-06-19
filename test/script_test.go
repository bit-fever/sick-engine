//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package test

import (
	"strings"
	"testing"

	"github.com/algotiqa/tiq-engine/core"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/tiq-engine/runtime"
)

//=============================================================================
//=== Package
//=============================================================================

func TestPackageNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("package A")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================
//=== Constants
//=============================================================================

func TestConstantNameInUppercase(t *testing.T) {
	_, pes := core.ParseString("const { a = 4 }")
	assertErrors(t, pes, "uppercase")
}

//=============================================================================

func TestConstantNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("const { A = 4 A = 5}")
	assertErrors(t, pes, "already used")
}

//=============================================================================
//=== Variables
//=============================================================================

func TestVariableNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("var { A = 4 }")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================

func TestVariableNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("var { a = 4 a = 5}")
	assertErrors(t, pes, "already used")
}

//=============================================================================
//=== Functions
//=============================================================================

func TestFunctionNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("func A() {}")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================

func TestFunctionClassNameInUppercase(t *testing.T) {
	_, pes := core.ParseString("func (b) a() {}")
	assertErrors(t, pes, "uppercase")
}

//=============================================================================

func TestFunctionNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("func a() {} func a() {}")
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestFunctionParamNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("func a(B int) {}")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================

func TestFunctionParamNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("func a(b int, b bool) {}")
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestFunctionCallIdentifierInLowercase(t *testing.T) {
	_, pes := core.ParseString("func a() { C() }")
	assertErrors(t, pes, "lowercase")

	_, pes = core.ParseString("func a() { b.C() }")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================
//=== Enums
//=============================================================================

func TestEnumNameInUppercase(t *testing.T) {
	_, pes := core.ParseString("enum a { B }")
	assertErrors(t, pes, "uppercase")
}

//=============================================================================

func TestEnumItemNameInUppercase(t *testing.T) {
	_, pes := core.ParseString("enum A { b }")
	assertErrors(t, pes, "uppercase")
}

//=============================================================================

func TestEnumItemNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("enum A { B B }")
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestEnumWithMixedItems(t *testing.T) {
	_, pes := core.ParseString("enum A { B(1) C(\"x\") }")
	assertErrors(t, pes, "mixing")
}

//=============================================================================

func TestEnumWithAutoCodeAssignment(t *testing.T) {
	s, pes := core.ParseString("enum A { B C }")
	if !pes.IsEmpty() {
		t.Errorf("Errors returned")
	}

	e := s.Enums[0]
	if e.Items[0].Code != 1 && e.Items[1].Code != 2 {
		t.Errorf("Bad auto assignment")
	}
}

//=============================================================================
//=== Classes
//=============================================================================

func TestClassNameInUppercase(t *testing.T) {
	_, pes := core.ParseString("class a { }")
	assertErrors(t, pes, "uppercase")
}

//=============================================================================

func TestPropertyNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("class A { B int }")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================

func TestClassNameDuplicated(t *testing.T) {
	_, pes := core.ParseString("class A {} class A { b int }")
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestClassPropertyDuplicated(t *testing.T) {
	_, pes := core.ParseString("class A { b int b bool }")
	assertErrors(t, pes, "already used")
}

//=============================================================================
//=== Identifiers
//=============================================================================

func TestIdentifierNameInLowercase(t *testing.T) {
	_, pes := core.ParseString("func a() { B=5 }")
	assertErrors(t, pes, "lowercase")

	_, pes = core.ParseString("func a() { a.B=5 }")
	assertErrors(t, pes, "lowercase")
}

//=============================================================================
//===
//=== Environment
//===
//=============================================================================

func TestConstantNameDuplicatedInFiles(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("const A=5")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("const A=7")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestVariableNameDuplicatedInFiles(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("var a=5")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("var a=7")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestFunctionNameDuplicatedInFiles(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("func a() {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("func a() {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestFunctionNameWithDifferentParameters(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("func a() {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("func a(b int) {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)
}

//=============================================================================

func TestFunctionNameWithDifferentScopes(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("func a() {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("func (C) a() {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)
}

//=============================================================================

func TestEnumNameDuplicatedInFiles(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("enum A { A }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("enum A { B }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestClassNameDuplicatedInFiles(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("class A {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("class A {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertErrors(t, pes, "already used")
}

//=============================================================================

func TestResolveEnumTypeNotExisting(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("class A { a Color }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("enum Size { X }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	err := e.Build()
	assertError(t, err, "resolve type")
}

//=============================================================================

func TestResolveEnumTypeExisting(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("class A { a Color }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("enum Color { X }")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	err := e.Build()
	assertNoError(t, err)
}

//=============================================================================

func TestResolveClassTypeNotExisting(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("func a(p Color) {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("class ColorX {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	err := e.Build()
	assertError(t, err, "resolve type")
}

//=============================================================================

func TestResolveClassTypeExisting(t *testing.T) {
	e := runtime.NewEnvironment()
	s, pes := core.ParseString("func a(p Color) {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	s, pes = core.ParseString("class Color {}")
	assertNoErrors(t, pes)
	pes = e.AddScript(s)
	assertNoErrors(t, pes)

	err := e.Build()
	assertNoError(t, err)
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func assertError(t *testing.T, err *parser.ParseError, expected string) {
	if err == nil {
		t.Errorf("No error returned")
	} else if !strings.Contains(err.Error, expected) {
		t.Errorf("Error should contain: %s", expected)
	}
}

//=============================================================================

func assertErrors(t *testing.T, pes *parser.ParseErrors, expected string) {
	if pes.IsEmpty() {
		t.Errorf("No error returned")
	} else if !strings.Contains(pes.Errors[0].Error, expected) {
		t.Errorf("Error should contain: %s", expected)
	}
}

//=============================================================================

func assertNoError(t *testing.T, err *parser.ParseError) {
	if err != nil {
		t.Errorf("Error returned. Expected none.")
	}
}

//=============================================================================

func assertNoErrors(t *testing.T, pes *parser.ParseErrors) {
	if !pes.IsEmpty() {
		t.Errorf("Errors returned. Expected none.")
	}
}

//=============================================================================
