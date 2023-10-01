package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"strcony"
)

func main() {

	//Create arrays
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	//perform OLS regression
	set_I := linear.NewRegression()
	set_I.PushAll(x1, y1)

	set_II := linear.NewRegression()
	set_II.PushAll(x2, y2)

	set_III := linear.NewRegression()
	set_III.PushAll(x3, y3)

	set_IV := linear.NewRegression()
	set_IV.PushAll(x4, y4)

	//print coefficients
	fmt.Println("Coefficients")
	set_I_coef_I, set_I_coef_II := set_I.Coefficients()
	fmt.Println("const set_I:", set_I_coef_I, "x1:", set_I_coef_II)

	set_II_coef_I, set_II_coef_II := set_II.Coefficients()
	fmt.Println("const set_II:", set_II_coef_I, "x2:", set_II_coef_II)

	set_III_coef_I, set_III_coef_II := set_III.Coefficients()
	fmt.Println("const set_III:", set_III_coef_I, "x3:", set_III_coef_II)

	set_IV_coef_I, set_IV_coef_II := set_IV.Coefficients()
	fmt.Println("const set_IV:", set_IV_coef_I, "x4:", set_IV_coef_II)
}
