
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