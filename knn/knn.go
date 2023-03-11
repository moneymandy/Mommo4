
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
}

func New(dataSet [][]float64, labels []string) KnnClassifier {
    kn := KnnClassifier{dataSet: dataSet, labels: labels}
    kn.calcMinMax()
    kn.normalizeData()