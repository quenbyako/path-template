// Copyright (c) 2021 Xelaj Software
//
// This file is a part of path-template package.
// See https://github.com/quenbyako/path-template/blob/master/LICENSE.md for details

package tpl_test

import (
	"fmt"

	tpl "github.com/quenbyako/path-template"
)

func ExampleFillTemplate() {
	fmt.Println(tpl.FillTemplate(
		"/join/{chat_id:[0-9a-f]+}/link",
		map[string]string{
			"chat_id": "12345abcdef",
		},
	))
	// Output: /join/12345abcdef/link nil
}

func ExampleFillTemplate_invalid() {
	// bitly link filler for example
	fmt.Println(tpl.FillTemplate(
		"http://bit.ly/{link_id}",
		map[string]string{},
	))
	// Output: error: "missed 'link_id' parameter"
}

func ExampleMatchPath() {
	fmt.Println(tpl.MatchPath("/join/{chat_id:[0-9a-f]+}/link", "/join/3298239fffadcf425/link"))
	// Output: map[string]string{chat_id: 3298239fffadcf425}
}

func ExampleGetTemplateVariables() {
	fmt.Println(tpl.GetTemplateVariables(
		"/etc/{package_name}/configs/{env_name}/.env",
	))
	// Output: [package_name env_name]
}
