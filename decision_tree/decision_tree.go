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
    for _, count := range labelCounts {
        prob := float64(count) / float64(numEntries)
        shannonEnt -= prob * math.Log2(prob)
    }
    // fmt.Println("Entropy: ", shannonEnt, "; dataSet: ", dataSet)
    return shannonEnt
}

func SplitDataSet(dataSet [][]int, axis int, value int) [][]int {
    var retDataSet [][]int
    // fmt.Println("splitting: ", dataSet)
    for _, featureVector := range dataSet {
        if featureVector[axis] == value {
            reducedFeatVector := make([]int, len(featureVector) - 1)
            copy(reducedFeatVector, featureVector[:axis])
            copy(reducedFeatVector[axis:], featureVector[axis+1:])
            retDataSet = append(retDataSet, reducedFeatVector)
        }
    }
    // fmt.Println("sp