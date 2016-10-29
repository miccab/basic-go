package algo

import (
	"testing"
	"github.com/assertgo/assert"
)

func TestFibonRecurZero(t *testing.T) {
	var a = assert.New(t)

	var result = FibonRecur(0)

	a.ThatInt(result).IsEqualTo(0)
}

func TestFibonRecurOne(t *testing.T) {
	var a = assert.New(t)

	var result = FibonRecur(1)

	a.ThatInt(result).IsEqualTo(1)
}


func TestFibonRecurThree(t *testing.T) {
	var a = assert.New(t)

	var result = FibonRecur(3)

	a.ThatInt(result).IsEqualTo(2)
}

func TestFibonRecurFive(t *testing.T) {
	var a = assert.New(t)

	var result = FibonRecur(5)

	a.ThatInt(result).IsEqualTo(5)
}
