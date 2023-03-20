
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
    return kn
}

func ( kn *KnnClassifier ) Classify( inX []float64, k int ) (string, error) {
    if len(kn.dataSet) == 0 {
        return "", errors.New("Empty data set")
    }
    diff := calcDiff(kn.normalizeInput(inX), kn.normalizedDataSet)
    sort.Sort(diff)
    // fmt.Println(diff)

    return kn.vote(diff, k), nil
}

func calcDiff(inX []float64, dataSet [][]float64) DiffArray{
    res := NewDiffArray(len(dataSet))
    // fmt.Println(res)
    for i, row := range dataSet {
        // fmt.Printf("row: %#v %d\n", row, i)
        for j := 0; j < len(inX); j++ {
            res.Values[i] += math.Pow(inX[j]*inX[j] - row[j]*row[j], 2)
        }
        res.Values[i] = math.Sqrt(res.Values[i])
    }
    return res
}

func ( kn KnnClassifier ) vote(diff DiffArray, k int) string {
    m := make(map[string]int)
    for i := 0; i < int(math.Min( float64(len(diff.Values)), float64(k)) ); i++ {
        m[kn.labels[diff.Indices[i]]] += 1
    }

    var max int
    var res string
    for k, v := range m {
        if v > max {
            max = v
            res = k
        }
    }

    return res
}

func ( kn *KnnClassifier ) normalizeData() {