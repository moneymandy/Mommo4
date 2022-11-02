package decisionTree

type Tree struct {
    SubTrees map[string]Tree
    Label string
    Feature string
}

func NewTree() Tree {
    tree :