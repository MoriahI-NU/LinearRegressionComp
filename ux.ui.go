package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inancgumus/screen"
)

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func mainMenu() {
	clearScreen()
	var choice int64 = -1
	var set string
	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Regression Performance - Main Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Calculate Performance Using Anscombe Quartet Set 1")
		fmt.Println("2. Calculate Performance Using Anscombe Quartet Set 2")
		fmt.Println("3. Calculate Performance Using Anscombe Quartet Set 3")
		fmt.Println("4. Calculate Performance Using Anscombe Quartet Set 4")
		fmt.Println("5. Calculate Performance Using ALL Anscombe Quartet Sets")
		fmt.Println("6. View Average Performance for All Tests in Current Session")
		fmt.Println("7. Graph Average Performance for All Tests in Current Session")
		fmt.Println("8. Help")
		fmt.Println("9. Configuration Menu")
		fmt.Println("0. Exit")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			fmt.Println("Goodbye...")
		case 1:
			clearScreen()
			set = "One"
			experiment(set)
			menuReset()
		case 2:
			clearScreen()
			set = "Two"
			experiment(set)
			menuReset()
		case 3:
			clearScreen()
			set = "Three"
			experiment(set)
			menuReset()
		case 4:
			clearScreen()
			set = "Four"
			experiment(set)
			menuReset()
		case 5:
			clearScreen()
			fmt.Println("running all Four Anscombe Quartets.")
			experiment("One")
			experiment("Two")
			experiment("Three")
			experiment("Four")
			menuReset()
		case 6:
			clearScreen()
			fmt.Println("Calculating Performance on all runs:")
			calcPerformance()
			menuReset()
		case 7:
			clearScreen()
			fmt.Println("Graphing Performance on all runs:")
			performanceGraph()
			menuReset()
		case 8:
			clearScreen()
			help()
			menuReset()
		case 9:
			clearScreen()
			optionsMenu()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		if choice == 0 {
			clearScreen()
			return
		}
	}
}

func optionsMenu() {
	clearScreen()
	fmt.Println("Configuration Menu")
	fmt.Println()

	//determine which set to test on

	var choice int64 = -1

	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Configuration Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Set number of runs per test")
		fmt.Println("2. Set nubmer of runs per experiment")
		fmt.Println("3. Set config to Debug mode")
		fmt.Println("4. Set rounding value of coefficients")
		fmt.Println("5. Set rounding value of runtime")
		fmt.Println("Press Enter to Return to Previous Menu")

		var err error
		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			clearScreen()
			fmt.Println()
		case 1:
			clearScreen()
			fmt.Println("Runs pertest determines number of times each model will be built.")
			fmt.Println("Current number of runs per test:", nRuns)
			fmt.Println("Please enter new number and press enter")
			_, err := fmt.Scanf("%d\n", &nRuns)
			if err != nil {
				fmt.Println("Invalid input. Please try again")
				continue
			}
			fmt.Println("Current number of runs per test:", nRuns)
			menuReset()
		case 2:
			clearScreen()
			fmt.Println("Runs per experiment determines number of times each experiment will be run")
			fmt.Println("Current number of runs per experiment:", nExperiments)
			fmt.Println("Please enter new number and press enter")
			_, err := fmt.Scanf("%d\n", &nExperiments)
			if err != nil {
				fmt.Println("Invalid input. Please try again")
				continue
			}
		case 3:
			clearScreen()
			fmt.Println("Setting config to debug mode")
			nExperiments = 1
			menuReset()
		case 4:
			clearScreen()
			fmt.Println("Rounding Value of Coefficients")
			fmt.Println("Current rounding value:", roundCoefficients)
			fmt.Println("Please enter new number and press enter")
			_, err := fmt.Scanf("%d\n", &roundCoefficients)
			if err != nil {
				fmt.Println("Invalid input. Please try again")
				continue
			}
			fmt.Println("Current rounding of Coefficients:", roundCoefficients)
			menuReset()
		case 5:
			clearScreen()
			fmt.Println("Rounding value of runtime")
			fmt.Println("Current rounding value:", roundTime)
			fmt.Println("Please enter new number and press enter")
			_, err := fmt.Scanf("%d\n", &roundTime)
			if err != nil {
				fmt.Println("Invalid input. Please try again")
				continue
			}
			fmt.Println("Curren trounding value of Runtime:", roundTime)
			menuReset()
		default:
			clearScreen()
			return
		}

		//if choice == 0
		//fmt.Println("Press Enter to continue...")
		//scanner.Scan() // Wait for user to press Enter

	}
}
