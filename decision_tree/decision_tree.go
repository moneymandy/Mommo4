package decisionTree

import (
    "fmt"
    "math"
)

func CalcShannonEntropy(dataSet [][]int) float64 {
    numEntries := len(dataSet)
    labelCounts :=