package main

import "C"
import (
	"fmt"
	"strings"
	"text/template"
)

// Functions returns a list of functions for jstpl
func Functions() template.FuncMap {
	return template.FuncMap{
		"strings":     toStrings,
		"contains":    strings.Contains,
		"count":       strings.Count,
		"equalFold":   strings.EqualFold,
		"fields":      strings.Fields,
		"hasPrefix":   strings.HasPrefix,
		"hasSuffix":   strings.HasSuffix,
		"index":       strings.Index,
		"join":        strings.Join,
		"lastIndex":   strings.LastIndex,
		"repeat":      strings.Repeat,
		"replace":     strings.Replace,
		"split":       strings.Split,
		"splitAfter":  strings.SplitAfter,
		"splitAfterN": strings.SplitAfterN,
		"splitN":      strings.SplitN,
		"title":       strings.ToTitle,
		"lower":       strings.ToLower,
		"upper":       strings.ToUpper,
		"trim":        strings.Trim,
		"trimLeft":    strings.TrimLeft,
		"trimPrefix":  strings.TrimPrefix,
		"trimRight":   strings.TrimRight,
		"trimSpace":   strings.TrimSpace,
		"trimSuffix":  strings.TrimSuffix,
	}
}

func toStrings(vs []interface{}) []string {
	var strs []string

	for _, v := range vs {
		strs = append(strs, fmt.Sprintf("%v", v))
	}

	return strs
}
