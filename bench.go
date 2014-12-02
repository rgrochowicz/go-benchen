package main

import (
	"testing"
	"encoding/csv"
	"strconv"
	"io"
	"log"
	"runtime"
)

type NCollection struct {
	N int
	Results []testing.BenchmarkResult
}

func (c *NCollection) ToStringSlice() []string {
	result := []string{
		strconv.Itoa(c.N),
	}

	for _, nres := range c.Results {
		result = append(result, strconv.Itoa(int(nres.NsPerOp())))
	}

	return result
}

type BenchFunc func(n int, b *testing.B)

type Benchem struct {
	F BenchFunc
	Name string
}

func NewBenchem(name string, f BenchFunc) *Benchem {
	return &Benchem{
		F: f,
		Name: name,
	}
}

func (bm *Benchem) Run(n int) testing.BenchmarkResult {
	return testing.Benchmark(func(b *testing.B) {
		bm.F(n, b)
	})
}

type BenchCollection struct {
	Benches []*Benchem
	Results []*NCollection
}

func NewBenchCollection() *BenchCollection {
	return &BenchCollection {
		Benches: []*Benchem{},
		Results: []*NCollection{},
	}
}

func (c *BenchCollection) AddFunc(name string, f BenchFunc) {
	c.Benches = append(c.Benches, NewBenchem(name, f))
}

func (c *BenchCollection) Run(start, end, step int) {

	c.Results = []*NCollection{}

	for iterations := start; iterations <= end; iterations += step {

		collection := &NCollection{
			N: iterations,
			Results: []testing.BenchmarkResult{},
		}
		c.Results = append(c.Results, collection)

	}

	for _, bench := range c.Benches {
		for _, collection := range c.Results {
			collection.Results = append(collection.Results, bench.Run(collection.N))

			runtime.GC()

			log.Printf("Completed run for %d iterations", collection.N)
		}

		log.Printf("Completed bench: %s", bench.Name)
	}

}

func (c *BenchCollection) BenchNames() []string {
	result := []string{}

	for _, bench := range c.Benches {
		result = append(result, bench.Name)
	}

	return result
}

func (c *BenchCollection) Save(w io.Writer) {
	writer := csv.NewWriter(w)

	csvTitles := []string{"N"}
	csvTitles = append(csvTitles, c.BenchNames()...)

	writer.Write(csvTitles)

	for _, result := range c.Results {
		writer.Write(result.ToStringSlice())
	}

	writer.Flush()
}