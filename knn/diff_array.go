
package knn

type DiffArray struct {
    Values []float64
    Indices []int
}

func NewDiffArray(len int) DiffArray {