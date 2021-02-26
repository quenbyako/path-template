// Copyright (c) 2021 Xelaj Software
//
// This file is a part of path-template package.
// See https://github.com/quenbyako/path-template/blob/master/LICENSE.md for details

package tpl_test

import (
	"fmt"

	tpl "github.com/quenbyako/path-template"
)

func ExampleGetTemplateVariables() {
	fmt.Println(tpl.GetTemplateVariables(
		"/etc/{package_name}/configs/{env_name}/.env",
	))
	// Output: [package_name env_name]
}
