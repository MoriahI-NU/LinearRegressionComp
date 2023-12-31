package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/guptarohit/asciigraph"
	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
)

// mean runtimes
func calcPerformance() {
	if len(performanceGo) == 0 {
		fmt.Println("No experiments have been run in this session")
		runall()
		//calcPerformance()
		//return
	}

	meanPython, _ := stats.Mean(performancePython)
	meanR, _ := stats.Mean(performanceR)
	meanGo, _ := stats.Mean(performanceGo)
	fmt.Println()
	fmt.Println("Mean Python Runtime:", fmt.Sprintf("%v", meanPython))
	fmt.Println("Mean R Runtime:", fmt.Sprintf("%v", meanR))
	fmt.Println("Mean Go Runtime:", fmt.Sprintf("%v", meanGo))
	fmt.Println()
	performanceMatrix(meanPython, meanR, meanGo)
}

// Create table of results
func createTable(resultsGo, resultsPython, resultsR Response) {
	data := [][]interface{}{
		{"Go",
			roundFloat(resultsGo.Coefficients[0], roundCoefficients),
			roundFloat(resultsGo.Coefficients[1], roundCoefficients),
			roundFloat(resultsGo.Time, roundTime),
		},
		{"Python",
			roundFloat(resultsPython.Coefficients[0], roundCoefficients),
			roundFloat(resultsPython.Coefficients[1], roundCoefficients),
			roundFloat(resultsPython.Time, roundTime),
		},
		{"R",
			roundFloat(resultsR.Coefficients[0], roundCoefficients),
			roundFloat(resultsR.Coefficients[1], roundCoefficients),
			roundFloat(resultsR.Time, roundTime),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Language", "Intersect", "Slope", "Runtime"})

	table.SetAutoWrapText(false)

	for _, row := range data {
		strRow := make([]string, len(row))
		for i, val := range row {
			if floatValue, ok := val.(float64); ok {
				strRow[i] = strconv.FormatFloat(floatValue, 'f', -1, 64)
			} else {
				strRow[i] = fmt.Sprint(val)
			}
		}
		table.SetBorder(false)
		table.Append(strRow)
	}

	table.Render()
	fmt.Println()

	return
}

// Speed Ratios
func performanceMatrix(meanPython, meanR, meanGo float64) {
	if len(performanceGo) == 0 {
		fmt.Println("No experiments have been run in this session")
		runall()
		//calcPerformance()
		//return
	}

	data := [][]interface{}{
		{"Go", meanGo / meanGo, meanPython / meanGo, meanR / meanGo},
		{"Python", meanGo / meanPython, meanPython / meanPython, meanR / meanPython},
		{"R", meanGo / meanR, meanPython / meanR, meanR / meanR},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Baseline", "Go (Numerator)", "Python (Numerator)", "R (Numerator)"})

	table.SetAutoWrapText(false)

	for _, row := range data {
		strRow := make([]string, len(row))
		for i, val := range row {
			if floatValue, ok := val.(float64); ok {
				strRow[i] = strconv.FormatFloat(floatValue, 'f', -1, 64)
			} else {
				strRow[i] = fmt.Sprint(val)
			}
		}
		table.SetBorder(false)
		table.Append(strRow)
	}

	table.Render()
	if len(performanceGo) == 0 {
		fmt.Println()
		fmt.Println("No experiments have been run in this session")
		return
	}
	fmt.Println()

	return
}

// Performance Graph
func performanceGraph() {

	if len(performanceGo) == 0 {
		fmt.Println("No experiments have been run this sesssion")
		runall()
	}
	log := [][]float64{
		performanceGo,
		performanceR,
		performancePython,
	}

	graph := asciigraph.PlotMany(log, asciigraph.Precision(10), asciigraph.SeriesColors(
		asciigraph.Red,
		asciigraph.Yellow,
		asciigraph.Green,
	))

	fmt.Println("Legend: Go (red), Python (green), R (yellow)")
	fmt.Println(graph)
	fmt.Println()
}
