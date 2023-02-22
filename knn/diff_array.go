
package knn

type DiffArray struct {
    Values []float64
    Indices []int
}

func NewDiffArray(len int) DiffArray {
    da := DiffArray{
        make([]float64, len),
        make([]int, len),
    }
    for i := 0; i < len; i++ {
        da.Indices[i] = i;
    }
    return da
}

func (da DiffArray) Len() int {
    return len(da.Values)