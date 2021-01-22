package main

import (
	"fmt"
	"github.com/codonex/towersampling/pkg/towersampling"
)

func main() {
	sampler := towersampling.Sampler([]towersampling.Action{
		{Name: "Sleep", Probability: 0.33},
		{Name: "Walk", Probability: 0.05},
		{Name: "Watch", Probability: 0.35},
		{Name: "Eat", Probability: 0.27},
	})

	action_freq := map[string]float32{}
	const N = 1000000
	for i := 0; i < N; i++ {
		value, _ := sampler()
		action_freq[value] += 1
	}

	for k, v := range action_freq {
		action_freq[k] = v / N
	}
	fmt.Println(action_freq)

}
