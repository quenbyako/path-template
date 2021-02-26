// Copyright (c) 2021 Xelaj Software
//
// This file is a part of path-template package.
// See https://github.com/quenbyako/path-template/blob/master/LICENSE.md for details

package tpl

import "strings"

func stringStringMapKeys(l map[string]string) []string {
	res := make([]string, 0, len(l))
	for k := range l {
		res = append(res, k)
	}
	return res
}

func stringsTrimSuffixPrefix(prefix, str, suffix string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, prefix), suffix)
}
