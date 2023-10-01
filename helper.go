package main

import "github.com/montanaflynn/stats"

func MakeCoordinates(x, y []float64) []stats.Coordinate {
	container := make([]stats.Coordinate, len(x))

	for i := 0; i < len(x); i++ {
		container[i] = stats.Coordinate{
			x[i], y[i],
		}
	}
	return container
}
