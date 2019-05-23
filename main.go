package main

import (
	"fmt"
	"log"

	"github.com/austingebauer/go-ray-tracer/tuple"
)

func main() {
	vec1 := tuple.NewVector(1, -2, 4)
	vec2 := tuple.NewVector(1, -2, 4)
	vecA, err := tuple.Add(vec1, vec2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", vecA)
	fmt.Printf("%v\n", tuple.Equals(vec1, vec2))
}
