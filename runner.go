package main

import (
	"io"
)

var DefaultBenchCollection = NewBenchCollection()

func AddFunc(name string, f BenchFunc) {
	DefaultBenchCollection.AddFunc(name, f)
}

func Run(start, end, step int) {
	DefaultBenchCollection.Run(start, end, step)
}

func Save(w io.Writer) {
	DefaultBenchCollection.Save(w)
}