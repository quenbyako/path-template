// Copyright (c) 2021 Xelaj Software
//
// This file is a part of path-template package.
// See https://github.com/quenbyako/path-template/blob/master/LICENSE.md for details

package tpl

import (
	"fmt"
	"regexp"
	"strings"
)

// MatchPath extracting path variables with template.
// it returns nil, false, if path doesn't match template
// got example: matchPath("/joinchat/{chat_id}", "/joinchat/abcdefg") returns {"chat_id":"abcdefg"}, false
// spiced up implementaition from https://git.io/Jtcv0 (cuz why not?)
func MatchPath(tpl, path string) (map[string]string, bool) {
	// if template doesn't have pattern
	if !strings.ContainsAny(tpl, "{}") {
		if tpl == path {
			return map[string]string{}, true
		}
		return nil, false
	}

	tplIsGlobal := strings.HasPrefix(tpl, "/")
	pathIsGlobal := strings.HasPrefix(path, "/")

	if tplIsGlobal {
		if !pathIsGlobal {
			return nil, false
		}

		tpl = strings.TrimLeft(tpl, "/")
	}
	if pathIsGlobal {
		if !tplIsGlobal {
			return nil, false
		}

		path = strings.TrimLeft(path, "/")
	}

	tplPathItems := strings.Split(tpl, "/")
	pathItems := strings.Split(path, "/")
	if len(tplPathItems) != len(pathItems) {
		return nil, false
	}

	res := make(map[string]string)
	for i, tplPathItem := range tplPathItems {
		// if this item not a variable, we just need to check it but don't extract
		if !strings.HasPrefix(tplPathItem, "{") || !strings.HasSuffix(tplPathItem, "}") {
			if tplPathItem != pathItems[i] {
				return nil, false
			}
			continue
		}

		// {chat_id} -> chat_id
		templateObject := stringsTrimSuffixPrefix("{", tplPathItem, "}")
		name := templateObject
		result := pathItems[i]

		if strings.Contains(templateObject, ":") {
			i := strings.Index(templateObject, ":")
			name = templateObject[:i]
			regexpr := templateObject[i+1:]
			if regexpr == "" {
				res[name] = result
				continue
			}

			// здесь мы просто точно уверены, что регекс не пустой
			// добавляем ^ и $ при необходимости, т.к. матчимся целиком по строчке
			if []rune(regexpr)[0] != '^' {
				regexpr = "^" + regexpr
			}
			if []rune(regexpr)[len(regexpr)-1] != '$' {
				regexpr += "$"
			}

			if !regexp.MustCompilePOSIX(regexpr).MatchString(result) {
				return nil, false
			}
		}

		res[name] = result
	}

	return res, true
}

func FillTemplate(tpl string, data map[string]string) (string, error) {
	// if template doesn't have pattern
	if !strings.ContainsAny(tpl, "{}") {
		if len(data) > 0 {
			return "", fmt.Errorf("unused keys: [%v]", strings.Join(stringStringMapKeys(data), ", "))
		}
		return tpl, nil
	}

	d := make(map[string]string)
	for k, v := range data {
		d[k] = v
	}

	var isAbstract bool
	// if template or path are not global filepath
	if !strings.HasPrefix(tpl, "/") {
		isAbstract = true
	}
	tplPathItems := strings.Split(tpl, "/")
	for i, tplPathItem := range tplPathItems {
		if !strings.HasPrefix(tplPathItem, "{") || !strings.HasSuffix(tplPathItem, "}") {
			continue
		}

		// {chat_id} -> chat_id
		dataKey := stringsTrimSuffixPrefix("{", tplPathItem, "}")
		v, ok := d[dataKey]
		if !ok {
			return "", fmt.Errorf("key '%v' not found", dataKey)
		}
		tplPathItems[i] = v

		delete(d, dataKey)
	}

	res := strings.Join(tplPathItems, "/")
	if !isAbstract {
		res = "/" + res
	}

	if len(d) > 0 {
		return "", fmt.Errorf("unused keys: [%v]", strings.Join(stringStringMapKeys(d), ", "))
	}

	return res, nil
}

func GetTemplateVariables(tpl string) []string {
	if !strings.ContainsAny(tpl, "{}") {
		return []string{}
	}

	if !strings.HasPrefix(tpl, "/") {
		tpl = strings.TrimLeft(tpl, "/")
	}

	parts := strings.Split(tpl, "/")

	// we divide capacity by 2 cuz most likely wildcards count is less than half of all path parts
	res := make([]string, 0, len(parts)/2) //nolint:gomnd // see note above
	for _, part := range parts {
		if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
			continue
		}

		part = stringsTrimSuffixPrefix("{", part, "}")
		res = append(res, part)
	}

	return res
}
