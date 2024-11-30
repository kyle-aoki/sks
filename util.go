package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	check(err)
	return t
}

func fromJson[T any](path string) T {
	var t T
	check(json.Unmarshal(must(os.ReadFile(path)), &t))
	return t
}

func toJson[T any](path string, t T) {
	bytes := must(json.MarshalIndent(t, "", "  "))
	check(os.WriteFile(path, bytes, 0600))
}

var tw = tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)

func tabprint(str string) {
	tw.Write([]byte(str))
}

func tabflush() {
	tw.Flush()
}

func exec(cond bool, fn func()) {
	if cond {
		fn()
		os.Exit(0)
	}
}

func panicRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
		os.Exit(1)
	}
}
