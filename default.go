package benchen

import (
	"io"
)

var DefaultBenchCollection = NewBenchCollection()

func AddFunc(name string, f BenchFunc) {
	DefaultBenchCollection.AddFunc(name, f)
}

func RunStep(start, end, step int) {
	DefaultBenchCollection.RunStep(start, end, step)
}

func Save(w io.Writer) {
	DefaultBenchCollection.Save(w)
}