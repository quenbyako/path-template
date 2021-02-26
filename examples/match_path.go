// Copyright (c) 2021 Xelaj Software
//
// This file is a part of path-template package.
// See https://github.com/quenbyako/path-template/blob/master/LICENSE.md for details

package tpl_test

import (
	"fmt"

	"github.com/quenbyako/path-template"
)

func ExampleMatchPath() {
	fmt.Println(tpl.MatchPath("/join/{chat_id:[0-9a-f]+}/link", "/join/3298239fffadcf425/link"))
	// Output: map[string]string{chat_id: 3298239fffadcf425}
}
