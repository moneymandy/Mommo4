
package bayes

import (
    // "fmt"
    "regexp"
    "math"
    "encoding/json"
    "io/ioutil"
    "strings"
)

var defaultThreshold int = 2
var defaultProb float64 = 0.25
var dataPath string = "."

type WordStat struct {
    Prob float64
    Occurrencies int
}

func (ws WordStat) CorrectedProb() float64 {
    if ws.Occurrencies < defaultThreshold {
        return defaultProb
    }
    if ws.Prob == 1 {
        return 0.99
    }
    if ws.Prob == 0 {
        return 0.01
    }
    return ws.Prob
}

type WordsMap map[string]WordStat

type BayesClassifier struct {
    Words WordsMap
    DocNumber int
    DocFrequency float64
    Name string
}

func NewBayesClassifier(name string) BayesClassifier {
    return BayesClassifier{
        Name: name,
        Words: make(WordsMap),
    }
}

func (b *BayesClassifier) Learn(text string, good bool) {