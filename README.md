# Tower sampling

Tower sampling samples events according to given discrete distribution.
It can be used to generate events for giving distribution.
For example, for following distribution,
```golang

const (
	Sleep = towersampling.Action("Sleep")
	Walk  = towersampling.Action("Walk")
	Watch = towersampling.Action("Watch")
	Eat   = towersampling.Action("Eat")
)

sampler := towersampling.Sampler([]towersampling.Event{
	{Name: Sleep, Probability: 0.33},
	{Name: Walk, Probability: 0.05},
	{Name: Watch, Probability: 0.35},
	{Name: Eat, Probability: 0.27},
})

value, err := sampler()
```


the sampler will return Sleep event with 33% probability and Walk event with 5% probability.
