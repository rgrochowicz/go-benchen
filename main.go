package main

import (
	"log"
	"testing"
	"os"
)

func main() {
	
	AddFunc("Slice Append", func (n int, b *testing.B) {


		for i := 0; i < b.N; i++ {

			slice := []int{}
			for j := 0; j < n; j++ {
				slice = append(slice, 1)
			}

		}

	})

	AddFunc("Array Set", func (n int, b *testing.B) {

		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, n)
			for j := 0; j < n; j++ {
				arr = append(arr, 1)
			}
		}

	})


	Run(0, 1000, 5)

	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	Save(f)


}