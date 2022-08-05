package towersampling

import (
	"errors"
	"math/rand"
)

// Action or event name
type Action interface{}

// Event is event object
type Event struct {
	Name        Action
	Probability float32
	Description string
}

// Sampler samples data according to given discrete distribution
// This can be used to generate events for giving distribution
// For example, for following distribution,
//	const (
// 		Sleep = towersampling.Action("Sleep")
// 		Walk  = towersampling.Action("Walk")
// 		Watch = towersampling.Action("Watch")
// 		Eat   = towersampling.Action("Eat")
// 	)
//	...
// 	sampler := towersampling.Sampler([]towersampling.Event{
// 		{Name: Sleep, Probability: 0.33},
// 		{Name: Walk, Probability: 0.05},
// 		{Name: Watch, Probability: 0.35},
// 		{Name: Eat, Probability: 0.27},
// 	})
// 	value, err := sampler()
// the sampler will return Sleep with 33% probability.
func Sampler(events []Event) func() (Action, error) {
	distributions := make([]float32, len(events))
	eventNames := make([]Action, len(events))
	for i, e := range events {
		if i == 0 {
			distributions[i] = e.Probability
		} else {
			distributions[i] = e.Probability + distributions[i-1]
		}
		eventNames[i] = e.Name
	}

	return func() (Action, error) {
		value := rand.Float32()
		for i, p := range distributions {
			if value <= p {
				return eventNames[i], nil
			}
		}
		return "", errors.New("no action")
	}

}
