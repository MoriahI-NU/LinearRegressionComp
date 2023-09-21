package main

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	df := dataframe.New(
		series.New([]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, series.Float, "x1"),
		series.New([]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, series.Float, "x2"),
		series.New([]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, series.Float, "x3"),
		series.New([]float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}, series.Float, "x4"),
		series.New([]float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}, series.Float, "y1"),
		series.New([]float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}, series.Float, "y2"),
		series.New([]float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}, series.Float, "y3"),
		series.New([]float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}, series.Float, "y4"),
	)

	fmt.Println(df)
}