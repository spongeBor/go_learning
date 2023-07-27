package client

import (
	"go_learning/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	e := series.GetFibonacciSeries(5)
	t.Log(e)
}
