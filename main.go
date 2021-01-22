package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Action Action object
type Action struct {
	Name        string
	Probability float32
}

// Sampler samples data according to the probability
//
func Sampler(actions []Action) func() (string, error) {
	probabilities := make([]float32, len(actions))
	events := make([]string, len(actions))
	for i, e := range actions {
		if i == 0 {
			probabilities[i] = e.Probability
		} else {
			probabilities[i] = e.Probability + probabilities[i-1]
		}
		events[i] = e.Name
	}

	return func() (string, error) {
		value := rand.Float32()
		for i, p := range probabilities {
			if value <= p {
				return events[i], nil
			}
		}
		return "", errors.New("No action")
	}

}

func main() {
	sampler := Sampler([]Action{
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
