package main

import (
	"math"
	"testing"
)

// Tolerance for floating point differences
func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

var tolerance = 1e-13

// Python Tests
// // Run on first anscombe set
func TestExperimentPythonOne(t *testing.T) {
	t.Parallel()
	var set = "One"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	//Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
	//Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
}

// // Run on second anscombe set
func TestExperimentPythonTwo(t *testing.T) {
	t.Parallel()
	var set = "Two"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	//Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
	//Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
}

// // Run on third anscombe set
func TestExperimentPythonThree(t *testing.T) {
	t.Parallel()
	var set = "Three"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	//Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
	//Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
}

// // Run on fourth anscombe set
func TestExperimentPythonFour(t *testing.T) {
	t.Parallel()
	var set = "Four"
	var got = ExperimentGo(set, 1).Coefficients
	var want = ExperimentPython(set, "1").Coefficients

	//Test Intersect
	if !withinTolerance(want[0], got[0], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
	//Test Slope
	if !withinTolerance(want[1], got[1], tolerance) {
		t.Errorf("want #{want}, got #{got}")
	}
}
