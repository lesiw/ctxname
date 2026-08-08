package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/ctxname"
)

func main() { singlechecker.Main(ctxname.Analyzer) }
