package towersampling

import (
	"errors"
	"math/rand"
)

// Action Action object
type Action struct {
	Name        string
	Probability float32
	Description string
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
