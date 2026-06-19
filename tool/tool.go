//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package tool

import (
	"os"
	"path/filepath"
	"strings"
)

//=============================================================================

func StartsWithLowerCase(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}

//=============================================================================

func StartsWithUpperCase(s string) bool {
	return s[0] >= 'A' && s[0] <= 'Z'
}

//=============================================================================

func FindFiles(path string, suffix string) ([]string, error) {
	var list []string

	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file == path {
			return nil
		}

		if !info.IsDir() {
			if strings.HasSuffix(file, suffix) {
				list = append(list, file)
			}
		}

		return nil
	})

	return list,err
}

//=============================================================================
