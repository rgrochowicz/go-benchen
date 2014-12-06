package main

import (
	"log"
	"testing"
	"os"
	"flag"
	"github.com/rgrochowicz/go-benchen"
)

func main() {

	flag.Parse()

	onesArray := make([]int, 1000)
	for j := 0; j < 1000; j++ {
		onesArray[j] = 1
	}
	
	benchen.AddFunc("Slice Append", func (n int, b *testing.B) {


		for i := 0; i < b.N; i++ {

			slice := []int{}
			for j := 0; j < n; j++ {
				slice = append(slice, 1)
			}

		}

	})

	benchen.AddFunc("Slice Append With Capacity", func (n int, b *testing.B) {

		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, n)
			for j := 0; j < n; j++ {
				arr = append(arr, 1)
			}
		}

	})

	benchen.AddFunc("Slice Append With Num", func (n int, b *testing.B) {

		for i := 0; i < b.N; i++ {
			arr := make([]int, n)
			for j := 0; j < n; j++ {
				arr[j] = 1
			}
		}

	})

	benchen.AddFunc("Slice Copy from Array", func (n int, b *testing.B) {

		for i := 0; i < b.N; i++ {
			arr := make([]int, n)
			copy(arr, onesArray)
		}

	})

	benchen.RunStep(0, 1000, 1)

	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	benchen.Save(f)


}