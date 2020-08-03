package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	var w []string
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		w = append(w, temp)
	}

	sort.SliceStable(w, func(i, j int) bool {
		wi := strings.Split(w[i], "")
		wj := strings.Split(w[j], "")
		return comp(wi, wj)
	})
	sort.Strings(w)

	fmt.Println(strings.Join(w, ""))
}

func comp(s, t []string) bool {
	if comp1(s, t) || comp2(s, t) {
		return true
	} else {
		return false
	}
}

func comp1(s, t []string) bool {
	l := len(s)
	if len(s) < len(t) {
		l = len(t)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if s[j] == t[j] && s[i] < t[i] {
				return true
			}
		}
	}
	return false
}

func comp2(s, t []string) bool {
	sl := len(s)
	tl := len(t)

	l := sl
	if sl > tl {
		l = tl
	}

	for i := 0; i < l; i++ {
		if s[i] == t[i] && sl < tl {
			return true
		}
	}
	return false
}
