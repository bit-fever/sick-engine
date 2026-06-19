//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

import (
	"github.com/algotiqa/tiq-engine/core/data"
	"github.com/algotiqa/tiq-engine/parser"
)

//=============================================================================
//===
//=== Types
//===
//=============================================================================

func ConvertType(tree parser.ITypeContext) Type {
	if tree.ListType() != nil {
		lt := tree.ListType()
		subType := ConvertType(lt.Type_())
		return NewListType(subType)
	}

	if tree.MapType() != nil {
		mt := tree.MapType()
		keyType := ConvertKeyType(mt.KeyType())
		valType := ConvertType(mt.Type_())
		return NewMapType(keyType, valType)
	}

	if tree.INT() != nil {
		return NewIntType()
	}

	if tree.REAL() != nil {
		return NewRealType()
	}

	if tree.BOOL() != nil {
		return NewBoolType()
	}

	if tree.STRING() != nil {
		return NewStringType()
	}

	if tree.TIME() != nil {
		return NewTimeType()
	}

	if tree.DATE() != nil {
		return NewDateType()
	}

	if tree.TimeSeriesType() != nil {
		tst := tree.TimeSeriesType()
		var subType Type = NewRealType()
		if tst.Type_() != nil {
			subType = ConvertType(tst.Type_())
		}
		return NewTimeSeriesType(subType)
	}

	if tree.ERROR() != nil {
		return NewErrorType()
	}

	fqi := data.NewFQIdentifier(tree.FqIdentifier())

	return NewToFindOutType(fqi, parser.NewInfo(tree))
}

//=============================================================================

func ConvertKeyType(tree parser.IKeyTypeContext) Type {
	if tree.INT() != nil {
		return NewIntType()
	}

	if tree.STRING() != nil {
		return NewStringType()
	}

	if tree.TIME() != nil {
		return NewTimeType()
	}

	if tree.DATE() != nil {
		return NewDateType()
	}

	panic("Unknown key type : " + tree.GetText())
	return nil
}

//=============================================================================
