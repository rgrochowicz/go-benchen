package benchen

import (
	"testing"
	"encoding/csv"
	"strconv"
	"io"
	"log"
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

func (c *BenchCollection) RunStep(start, end, step int) {
	iterations := []int{}

	for i := start; i <= end; i += step {
		iterations = append(iterations, i)
	}

	c.Run(iterations)

}

func (c *BenchCollection) Run(iterations []int) {
	c.Results = []*NCollection{}

	for _, iteration := range iterations {

		collection := &NCollection{
			N: iteration,
			Results: []testing.BenchmarkResult{},
		}
		c.Results = append(c.Results, collection)

	}

	c.run()
}

func (c *BenchCollection) run() {

	for _, bench := range c.Benches {
		for _, collection := range c.Results {

			collection.Results = append(collection.Results, bench.Run(collection.N))

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