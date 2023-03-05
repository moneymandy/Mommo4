
package knn

import (
    // "fmt"
    "errors"
    "math"
    "sort"
)

type KnnClassifier struct {
    dataSet, normalizedDataSet [][]float64
    mins, maxes []float64
    labels  []string