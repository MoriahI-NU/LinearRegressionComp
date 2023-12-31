package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/montanaflynn/stats"
)

// Run Go Experiment
func ExperimentGo(set string, nRuns int) Response {
	//fmt.Println(" Performing Go Experiment")

	var responseGo Response
	//var times []float64

	startTime := time.Now()
	time.Sleep(startTime.Truncate(time.Millisecond).Add(time.Millisecond).Sub(startTime))
	startTime = time.Now()
	for i := 0; i < nRuns; i++ {

		points, _ := stats.LinearRegression(
			MakeCoordinates(
				anscombe[set]["x"],
				anscombe[set]["y"]),
		)
		responseGo.Coefficients = EquationLine(points)

	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Seconds()
	elapsedTime = roundFloat(elapsedTime/float64(nRuns), 10)
	if elapsedTime == 0 {
		elapsedTime = 0.00000000001
	}
	responseGo.Time = elapsedTime

	return responseGo
}

// Run Python Experiment
func ExperimentPython(set, nRunsString string) Response {
	var responsePython Response

	//Run Script
	args := []string{"AnscombeTest.py", set, nRunsString}
	cmd := exec.Command("python", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error Python:", err)
	}

	//parse response
	err = json.Unmarshal([]byte(output), &responsePython)
	if err != nil {
		fmt.Println("Error Python Response:", err)
	}

	return responsePython
}

// Run R Experiment
func ExperimentR(set, nRunsString string) Response {
	var responseR Response

	//Run Script
	args := []string{"AnscombeTest.R", set, nRunsString}
	cmd := exec.Command("Rscript", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error R:", err)
	}

	//Parse response
	err = json.Unmarshal([]byte(output), &responseR)
	if err != nil {
		fmt.Println("Error R Response:", err)
	}

	return responseR
}
