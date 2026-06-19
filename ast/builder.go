//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package ast

import (
	"strconv"

	"github.com/algotiqa/tiq-engine/ast/expression"
	"github.com/algotiqa/tiq-engine/ast/statement"
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/core/types"
	"github.com/algotiqa/tiq-engine/parser"
	"github.com/algotiqa/tiq-engine/tool"
)

//=============================================================================

const (
	DefaultPackage = "main"
)

//=============================================================================
//===
//=== AST builder
//===
//=============================================================================

func Build(tree parser.IScriptContext) *Script {
	script := NewScript(tree.GetParser().GetInputStream().GetSourceName())

	script.PackageName = convertPackage(tree.PackageDef())

	convertConstants(script, tree.AllConstantsDef())
	convertVariables(script, tree.AllVariablesDef())
	convertFunctions(script, tree.AllFunctionDef())
	convertEnums(script, tree.AllEnumDef())
	convertClasses(script, tree.AllClassDef())

	return script
}

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

func convertPackage(tree parser.IPackageDefContext) string {
	if tree != nil {
		id := tree.IDENTIFIER().GetText()

		if !tool.StartsWithLowerCase(id) {
			parser.RaiseError(tree.GetParser(), "package's name must start with a lowercase character: "+id)
		}

		return id
	}

	return DefaultPackage
}

//=============================================================================
//=== Constants
//=============================================================================

func convertConstants(script *Script, tree []parser.IConstantsDefContext) {
	for _, cds := range tree {
		for _, cd := range cds.AllConstantDef() {
			c := convertConstantDef(cd)
			script.AddConstant(c)
		}
	}
}

//=============================================================================

func convertConstantDef(tree parser.IConstantDefContext) *Constant {
	name := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "constant's name must start with an uppercase character: "+name)
	}

	return NewConstant(name, expression.ConvertExpression(value), parser.NewInfo(tree))
}

//=============================================================================
//=== Variables
//=============================================================================

func convertVariables(script *Script, tree []parser.IVariablesDefContext) {
	for _, vds := range tree {
		for _, vd := range vds.AllVariableDef() {
			v := convertVariableDef(vd)
			script.AddVariable(v)
		}
	}
}

//=============================================================================

func convertVariableDef(tree parser.IVariableDefContext) *Variable {
	name := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "variable's name must start with a lowercase character: "+name)
	}

	return NewVariable(name, expression.ConvertExpression(value), parser.NewInfo(tree))
}

//=============================================================================
//=== Functions
//=============================================================================

func convertFunctions(script *Script, tree []parser.IFunctionDefContext) {
	for _, fd := range tree {
		f := convertFunctionDef(fd)
		script.AddFunction(f)
	}
}

//=============================================================================

func convertFunctionDef(tree parser.IFunctionDefContext) *types.Function {
	name := tree.IDENTIFIER().GetText()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "function's name must start with a lowercase character: "+name)
	}

	f := types.NewFunction(name, parser.NewInfo(tree))

	if tree.Class() != nil {
		fqi := data.NewFQIdentifier(tree.Class().FqIdentifier())
		if !tool.StartsWithUpperCase(fqi.Name) {
			parser.RaiseError(tree.GetParser(), "function's class name must start with an uppercase character: "+fqi.Name)
		}

		f.Class = fqi
	}

	convertFunctionParams(f, tree.Parameters())
	convertFunctionResults(f, tree.Results())
	f.Block = statement.ConvertBlock(tree.Block())

	return f
}

//=============================================================================

func convertFunctionParams(f *types.Function, params parser.IParametersContext) {
	for _, pd := range params.AllParameterDecl() {
		name := pd.IDENTIFIER().GetText()
		type_ := pd.Type_()

		if !tool.StartsWithLowerCase(name) {
			parser.RaiseError(params.GetParser(), "function's parameter must start with a lowercase character: "+name)
		}

		t := types.ConvertType(type_)
		p := types.NewParam(name, t, parser.NewInfo(pd))
		f.AddParam(p)
	}
}

//=============================================================================

func convertFunctionResults(f *types.Function, results parser.IResultsContext) {
	if results != nil {
		for _, r := range results.AllType_() {
			t := types.ConvertType(r)
			f.AddReturnType(t)
		}
	}
}

//=============================================================================
//=== Enums
//=============================================================================

func convertEnums(script *Script, tree []parser.IEnumDefContext) {
	for _, ed := range tree {
		e := convertEnumDef(ed)
		script.AddEnum(e)
	}
}

//=============================================================================

func convertEnumDef(tree parser.IEnumDefContext) *types.EnumType {
	name := tree.IDENTIFIER().GetText()
	e := types.NewEnumType(name, true, parser.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "enum's name must start with an uppercase character: "+name)
	}

	countInts := 0
	countStr := 0

	for _, ei := range tree.AllEnumItem() {
		i := convertEnumItem(ei)
		e.AddItem(i)

		if i.Code != -1 {
			countInts++
		}

		if i.Value != "" {
			countStr++
		}
	}

	badInts := countInts > 0 && countInts != e.Size()
	badStr := countStr > 0 && countStr != e.Size()

	if badInts && badStr {
		parser.RaiseError(tree.GetParser(), "mixing integers and strings in enum: "+name)
	}

	if countInts == 0 && countStr == 0 {
		e.AssignCodes()
	} else if countStr != 0 {
		e.IsInt = false
	}

	return e
}

//=============================================================================

func convertEnumItem(tree parser.IEnumItemContext) *types.EnumItem {
	name := tree.IDENTIFIER().GetText()
	code := -1
	value := ""

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "enum item's name must start with an uppercase character: "+name)
	}

	enumVal := tree.EnumValue()
	if enumVal != nil {
		intValue := enumVal.INT_VALUE()
		strValue := enumVal.STRING_VALUE()

		if intValue != nil {
			iValue, err := strconv.Atoi(intValue.GetText())
			if err != nil {
				parser.RaiseError(tree.GetParser(), "Invalid integer value for enum : "+intValue.GetText())
			}

			code = iValue
		}
		if strValue != nil {
			text := strValue.GetText()
			value = text[1 : len(text)-1]
		}
	}
	return types.NewEnumItem(name, code, value)
}

//=============================================================================
//=== Classes
//=============================================================================

func convertClasses(script *Script, tree []parser.IClassDefContext) {
	for _, cd := range tree {
		c := convertClassDef(cd)
		script.AddClass(c)
	}
}

//=============================================================================

func convertClassDef(tree parser.IClassDefContext) *types.ClassType {
	name := tree.IDENTIFIER().GetText()
	c := types.NewClassType(name, parser.NewInfo(tree))

	if !tool.StartsWithUpperCase(name) {
		parser.RaiseError(tree.GetParser(), "class's name must start with an uppercase character: "+name)
	}

	for _, p := range tree.AllProperty() {
		c.AddProperty(convertProperty(p))
	}

	return c
}

//=============================================================================

func convertProperty(tree parser.IPropertyContext) *types.Property {
	name := tree.IDENTIFIER().GetText()
	type_ := tree.Type_()

	if !tool.StartsWithLowerCase(name) {
		parser.RaiseError(tree.GetParser(), "class's property must start with a lowercase character: "+name)
	}

	t := types.ConvertType(type_)

	return types.NewProperty(name, t, parser.NewInfo(tree))
}

//=============================================================================
