package deepcopy_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/codingpot/the-petty/deepcopy"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func elementsAreEqual(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}

type emptyStruct struct {
	Field1 string
	Field2 []int
}

func emptyStructGen() gopter.Gen {
	return gen.Struct(reflect.TypeOf(&emptyStruct{}), map[string]gopter.Gen{
		"Field1": gen.AnyString(),
		"Field2": gen.SliceOf(gen.Int()),
	})
}

func TestCopyShouldCreateANewCopy(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("deepcopy should return the same type", prop.ForAll(
		func(x interface{}) bool {
			newCopy, err := deepcopy.Copy(x)
			if err != nil {
				return false
			}
			return reflect.TypeOf(newCopy) == reflect.TypeOf(x)
		}, gen.OneGenOf(
			gen.Int(),
			gen.Float64(),
			gen.SliceOf(gen.Int()),
			gen.SliceOf(emptyStructGen()),
		)))

	properties.Property("[]int deepcopy should allocate a new copy", prop.ForAll(
		func(xs []int) bool {
			_newCopy, err := deepcopy.Copy(xs)

			newCopy := _newCopy.([]int)

			if err != nil {
				return false
			}

			if !elementsAreEqual(newCopy, xs) {
				return false
			}

			if len(xs) > 0 {
				idx := rand.Int() % len(xs)
				newCopy[idx] += 1
				return newCopy[idx] != xs[idx]
			}

			return true
		},
		gen.SliceOf(gen.Int()),
	))

	properties.TestingRun(t)
}
