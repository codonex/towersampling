package main

import (
	"fmt"

	"github.com/codonex/towersampling"
)

func main() {
	sampler := towersampling.Sampler([]towersampling.Action{
		{Name: "Sleep", Probability: 0.33},
		{Name: "Walk", Probability: 0.05},
		{Name: "Watch", Probability: 0.35},
		{Name: "Eat", Probability: 0.27},
	})

	actionFreq := map[string]float32{}
	const N = 1000000
	for i := 0; i < N; i++ {
		value, _ := sampler()
		actionFreq[value]++
	}

	for k, v := range actionFreq {
		actionFreq[k] = v / N
	}
	fmt.Println(actionFreq)

}
