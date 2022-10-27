package decisionTree

type Tree struct {
    SubTrees map[string]Tree
    Label string
    Feat