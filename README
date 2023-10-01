<h2 align="center">
Linear Regression Comparison of the Anscombe Quartet between R, Python, and Go
</h2>

<h4 align="center">
Please note that all code for this project was re-created from work done by Andrew D'Amico.
Problem setup and conditions were created by Northwestern University SPS MSDS 431
</h4>


## The Problem
The managers of a hypothetical tech startup company would like to have all employees use the same coding language - GO. In order to convince the company's data scientists of Go's performance and accuracy, a comparison test will be performed. Results from a linear regression test on four small subsets of the Anscombe dataset will be compared between Python, R, and Go. Run times between each experiments will also be compared.

## Running the Experiments
This project was completed with an executable file. You may run "anscombe.exe" to perform all linear regression comparisons on your device. It will present you with options to run each individual subset of the Anscombe Data, to run  all subsets at once, to view and graph the average of all runs, to configure settings, and to provide help. 

Choosing an option that calculates the performance of a subset of data will automatically output a table comparing the coefficient and runtime results between all three languages.

The configuration menu allows you to change the number of runs per test, the number of runs per experiment, and gives you the ability to set the decimal output of coefficient and runtime results.

# Personal Results

## Subset Results
Configuration settings were set to the following
Runs per test: 500
Runs per experiment: 15

### Performance on Anscombe Quartet Set One:

  LANGUAGE | INTERSECT |  SLOPE   |  RUNTIME    
-----------+-----------+----------+-------------
  Go       |  3.000091 | 0.500091 | 0.00000091  
  Python   |  3.000091 | 0.500091 | 0.00020209  
  R        |  3.000091 | 0.500091 | 0.00067086  

### Performance on Anscombe Quartet Set Two:

  LANGUAGE | INTERSECT |   SLOPE   |  RUNTIME
-----------+-----------+-----------+-------------
  Go       | 3.0009091 |       0.5 | 0.00000035
  Python   | 3.0009091 | 0.5000001 | 0.00018959
  R        | 3.0009091 | 0.5000001 | 0.00066988

### Performance on Anscombe Quartet Set Three:

  LANGUAGE | INTERSECT |   SLOPE   |  RUNTIME
-----------+-----------+-----------+-------------
  Go       | 3.0024546 | 0.4997273 | 0.00000048
  Python   | 3.0024546 | 0.4997273 | 0.00018959
  R        | 3.0024546 | 0.4997273 | 0.00067272

### Performance on Anscombe Quartet Set Four:

  LANGUAGE | INTERSECT |   SLOPE   |  RUNTIME
-----------+-----------+-----------+-------------
  Go       | 3.0017273 | 0.4999091 | 0.00000055
  Python   | 3.0017273 | 0.4999091 | 0.00019167
  R        | 3.0017273 | 0.4999091 | 0.00067199


## Average Performances from all Tests Ran during Session

### Mean Runtimes:

Mean Python Runtime: 0.0001932291666666668
Mean R Runtime: 0.0006713579098383584
Mean Go Runtime: 5.696243333333337e-07


### Table of Average Run Time Ratios:

  BASELINE |    GO (NUMERATOR)     | PYTHON (NUMERATOR)  |   R (NUMERATOR)
-----------+-----------------------+---------------------+---------------------
  Go       |                     1 |  339.22210720164696 | 1178.5976661314644
  Python   | 0.0029479210781671156 |                   1 |  3.474412902667513
  R        | 0.0008484659597895869 | 0.28781841076869147 |                  1


### Graph of Average Run Time Performances

Legend: Go (red), Python (green), R (yellow)
 0.0007062821388 ┼───────────────────────────────────────────────────────────
 0.0006053846918 ┤
 0.0005044872449 ┤
 0.0004035897979 ┤
 0.0003026923509 ┤
 0.0002017949039 ┼───────────────────────────────────────────────────────────
 0.0001008974570 ┤
 0.0000000000100 ┼───────────────────────────────────────────────────────────

## Recommendations to Management

For three out of the four subset tests, Go was able to obtain the same coefficients for intercepts and slopes. The only experiment that did not produce exact results were in subset two where the slope was a mere 0.0000001 units off. This difference is too small to cause any real issues for the purposes of most statistical analyses. Overall, in terms of coefficient results, Go holds up when compared to R and Python. Both Python and R performed exactly the same when making these calculations.

The average run times for Go were consistently better than the other languages tested regardless of subset. You can see from the table of average run time performances that Go performed over 400 times better than Python and over 1,000 times better than R! The Graph that follows clearly shows that Go's superior runtimes are followed by Python, with R falling far behind both.

I would recommend to Management that their plans of having all employees use Go should continue. It will bring comparable results with increased speed along with a greater sense of consistency and compatability across the whole company. It is important to note, however, that switching over to a new language involves a learning curve. Productivity may slow down while the data scientists familiarize themselves with using Go for analysis. There may be pushback from those resistant to change.