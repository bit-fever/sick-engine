//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package main

import (
	"github.com/algotiqa/tiq-engine/core"
)

//=============================================================================

func main() {
	//filename := os.Args[1]
	//res := ParseFile(filename)
	//if res.Errors != nil {
	//	return
	//}
	//
	//println(res.Script)

	e, errs := core.CreateEnvironment("sample")
	if !errs.IsEmpty() {
		println(errs.String())
	} else {
		_ = e
		println("Environment created")
	}
}

//=============================================================================
