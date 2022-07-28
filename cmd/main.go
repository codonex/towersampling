package main

import (
	"fmt"

	"github.com/codonex/towersampling"
)

const (
	Sleep = towersampling.Action("Sleep")
	Walk  = towersampling.Action("Walk")
	Watch = towersampling.Action("Watch")
	Eat   = towersampling.Action("Eat")
)

func main() {
	sampler := towersampling.Sampler([]towersampling.Event{
		{Name: Sleep, Probability: 0.33},
		{Name: Walk, Probability: 0.05},
		{Name: Watch, Probability: 0.35},
		{Name: Eat, Probability: 0.27},
	})

	actionFreq := map[towersampling.Action]float32{}
	const N = 1000000
	for i := 0; i < N; i++ {
		value, err := sampler()
		if err != nil {
			panic("Failed to sample data...")
		}
		actionFreq[value]++
	}

	for k, v := range actionFreq {

		actionFreq[k] = v / N
	}
	fmt.Println(actionFreq)

}
