package decisionTree

import (
    "fmt"
    "math"
)

func CalcShannonEntropy(dataSet [][]int) float64 {
    numEntries := len(dataSet)
    labelCounts := make(map[int]int)
    for i := 0; i < numEntries; i++ {
        currentLabel := dataSet[i][len(dataSet[i]) - 1]
        labelCounts[currentLabel] += 1
    }

    shannonEnt := 0.0
    for _, count := range labelCou